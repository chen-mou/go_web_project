package user

import (
	"errors"
	"gorm.io/gorm"
	"project/main/module/user/entity"
	"project/main/tool/dbTool"
	"reflect"
)

var Policies = make(map[string]*entity.Role)

var cache = make(map[string]*entity.Role)

func Register(path, name string) {
	role := entity.Role{}
	if cache[name] != nil {
		Policies[path] = cache[name]
		return
	}
	err := dbTool.Mysql.Where(" name = ? and status = ?", name, "NORMAL").Find(&role).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(any("没有找到目标权限：" + name))
		}
		panic(any(err.Error()))
	}
	Policies[path] = &role
	cache[name] = Policies[path]
}

func RegisterByStruct(obj any) {
	t := reflect.TypeOf(obj)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Type.Kind().String() == "func" {
			path := field.Tag.Get("path")
			role := field.Tag.Get("role")
			Register(path, role)
		}
	}
}
