package server

import (
	"os"
	"project/main/module/user/entity"
	"project/main/module/user/model"
	"project/main/tool/dbTool"
	"project/main/tool/encryption"
	"project/main/tool/jwtTool"
	"strconv"
	"time"
)

func Login(name, password string) (string, string) {
	user, _ := model.GetByName(name)
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
	ok := dbTool.GetLoopLock(key,
		strconv.Itoa(os.Getpid()),
		time.Second*5, 3000)
	defer dbTool.Unlock(key,
		strconv.Itoa(os.Getpid()))
	if !ok {
		return nil, "服务器繁忙"
	}
	_, err := model.GetByName(name)
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
