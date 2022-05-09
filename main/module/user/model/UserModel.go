package model

import (
	"encoding/json"
	"fmt"
	"project/main/module/user/entity"
	"project/main/tool"
	"project/main/tool/dbTool"
	"project/main/tool/encryption"
	"project/main/tool/time"
	"strconv"
	local "time"
)

func GetBaseByName(name string) (*entity.User, error) {
	user := entity.User{}
	var in int64 = 123456789
	val, _ := json.Marshal(time.Timestamp{
		Val: &in,
	})
	fmt.Println(string(val))
	err := dbTool.Mysql.Where("name = ? and status = 'NORMAL'", name).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

//func GetBaseByUUID(UUID string) (*entity.User, error) {
//
//}

func Create(users []entity.User, roleId int) ([]entity.User, error) {
	unix := local.Now().Unix()
	for i := range users {
		user := &users[i]
		user.Status = "NORMAL"
		user.Salt = encryption.MD5Salt(user.Name, strconv.FormatInt(unix, 10))
		user.Password = encryption.MD5SaltCount(user.Password, user.Salt, 5)
		user.UUID = encryption.MD5Salt(user.Name,
			tool.Get("name")+strconv.FormatInt(unix, 10))
		user.Ctime = time.Timestamp{
			Val: &unix,
		}
		user.Roles = []entity.UserRole{
			{
				RoleId: roleId,
				UserId: user.UUID,
				Ctime: time.Timestamp{
					Val: &unix,
				},
			},
			{
				RoleId: 9,
				UserId: user.UUID,
				Ctime: time.Timestamp{
					Val: &unix,
				},
			},
		}
		user.Data = entity.UserData{
			Name:   user.Name,
			Avatar: "default",
			Ctime: time.Timestamp{
				Val: &unix,
			},
		}
	}
	err := dbTool.Mysql.Create(&users).Error
	return users, err
}

//func CreateManager(users []entity.User) ([]entity.User, error) {
//
//}
