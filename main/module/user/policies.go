package user

import (
	"project/main/module/user/entity"
	"project/main/tool/dbTool"
	"reflect"
)

var Policies map[string]*entity.Role

func Register(path, name string) {
	role := entity.Role{}
	err := dbTool.Mysql.Where("where name = ? and status = ?", name, "NORMAL").Find(&role).Error
	if err != nil {
		panic(any(err))
	}
	Policies[path] = &role
}

func RegisterByStruct(obj any) {
	t := reflect.TypeOf(obj)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Type.Kind().String() == "func" {
			path := field.Tag.Get("path")
			role := field.Tag.Get("roleName")
			Register(path, role)
		}
	}
}
