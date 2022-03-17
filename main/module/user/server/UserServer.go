package server

import (
	"project/main/module/user/entity"
	"project/main/module/user/model"
	"project/main/tool/dbTool"
	"project/main/tool/encryption"
	"project/main/tool/jwtTool"
	"time"
)

func Login(name, password string) (string, string) {
	user, _ := model.GetBaseByName(name)
	if user == nil {
		return "", "用户名不存在"
	}
	if user.Password == encryption.MD5SaltCount(password, user.Salt, 5) {
		return jwtTool.GetToken(user.UUID, nil), ""
	}
	return "", "密码有误"
}

func Register(name, password string) (*entity.User, string) {
	key := "USER_REGISTER_LOCK_" + name
	value := dbTool.GetThreadID()
	ok := dbTool.GetLoopLock(key, value,
		time.Second*5, 3000)
	if !ok {
		return nil, "服务器繁忙"
	}
	defer dbTool.Unlock(key, value)
	_, err := model.GetBaseByName(name)
	if err == nil {
		return nil, "用户名已存在"
	}
	user, err1 := model.Create(name, password)
	if err1 != nil {
		panic(any(err1))
	}
	user.Salt = ""
	user.Id = -1
	return user, ""
}
