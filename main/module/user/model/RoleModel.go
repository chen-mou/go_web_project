package model

import (
	"gorm.io/gorm"
	"math/rand"
	"project/main/module/user/entity"
	"project/main/tool/dbTool"
	"time"
)

func GetUserRoleByUUID(uuid string) ([]entity.UserRole, string) {
	var value []entity.UserRole
	key := "USER_ROLE_GET_UUID_" + uuid
	dbTool.Get(key, value)
	if value == nil {
		err := dbTool.Mysql.Where("user_id = ? and status = ?", uuid, "NORMAL").Find(&value).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				dbTool.Set(key, nil, 5*time.Second)
			}
			return nil, err.Error()
		}
		dbTool.Set(key, value, 20*time.Minute+time.Duration(rand.Intn(5))*time.Minute)
	}
	return value, ""
}

func Verify(role entity.Role, target entity.Role) bool {
	auth := make(map[string]int32)
	for i := range role.Auths {
		auth[role.Auths[i].Auth.Scope] = role.Auths[i].Auth.Auth
	}
	for j := range target.Auths {
		if (auth[target.Auths[j].Auth.Scope] &
			target.Auths[j].Auth.Auth) != auth[target.Auths[j].Auth.Scope] {
			return false
		}
	}
	return true
}

func Mix(role entity.Role) *entity.Role {
	scope := make(map[string]int32)
	auths := role.Auths
	for i := range auths {
		scope[auths[i].Auth.Scope] = scope[auths[i].Auth.Scope] | auths[i].Auth.Auth
	}
	res := make([]entity.RoleAuth, len(scope))
	i := 0
	for k, v := range scope {
		res[i] = entity.RoleAuth{
			Auth: entity.Auth{
				Scope: k,
				Auth:  v,
			},
		}
		i++
	}
	role.Auths = res
	return &role
}
