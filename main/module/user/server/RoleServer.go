package server

import (
	"errors"
	"gorm.io/gorm"
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

func CreateBaseManager(UUIDs []string, expireTime []int32, roleId int) ([]entity.UserRole, error) {
	res := make([]entity.UserRole, len(UUIDs))
	err := dbTool.Mysql.Transaction(func(tx *gorm.DB) error {
		value := dbTool.GetThreadID()
		for index := range UUIDs {
			key := "CREATE_BASE_MANAGER" + UUIDs[index]
			ok := dbTool.GetLoopLock(key, value,
				time.Second*5, 3000)
			if !ok {
				return errors.New("服务器繁忙")
			}
			defer dbTool.Unlock(key, value)
			ok1, err := model.HasRole(UUIDs[index], roleId)
			if err != nil {
				return err
			}
			if !ok1 {
				continue
			}
			if expireTime != nil {
				res[index], err = model.CreateRole(tx, roleId, UUIDs[index], expireTime[index])
			} else {
				res[index], err = model.CreateRole(tx, roleId, UUIDs[index], 0)
			}
			if err != nil {
				return nil
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
