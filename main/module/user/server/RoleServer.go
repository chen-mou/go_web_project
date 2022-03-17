package server

import (
	"project/main/module/user/entity"
	"project/main/module/user/model"
	"project/main/tool/dbTool"
	"time"
)

func GetUserRoleByUUID(uuid string) ([]entity.UserRole, string) {
	value := dbTool.GetThreadID()
	key := "USER_ROLE_READ_BY_UUID_" + uuid
	ok := dbTool.GetLoopLock(key,
		value,
		5*time.Second,
		1000)
	if !ok {
		return nil, "服务器繁忙"
	}
	defer dbTool.Unlock(key, value)
	return model.GetUserRoleByUUID(uuid)
}
