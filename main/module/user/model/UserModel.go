package model

import (
	"golang.org/x/sys/windows"
	"project/main/module/user/entity"
	"project/main/tool/dbTool"
	"project/main/tool/encryption"
	"project/main/tool/time"
	"strconv"
	local "time"
)

func GetBaseByName(name string) (*entity.User, error) {
	user := entity.User{}
	err := dbTool.Mysql.Where("name = ? and status = 'NORMAL'", name).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

//func GetBaseByUUID(UUID string) (*entity.User, error) {
//
//}

func Create(name, password string) (*entity.User, error) {
	user := entity.User{}
	unix := local.Now().Unix()
	user.Name = name
	user.Status = "NORMAL"
	user.Salt = encryption.MD5Salt(user.Name, strconv.FormatInt(unix, 10))
	user.Password = encryption.MD5SaltCount(password, user.Salt, 5)
	user.UUID = encryption.MD5Salt(name,
		"machine_name"+strconv.FormatInt(unix, 10)+
			strconv.FormatUint(uint64(windows.GetCurrentThreadId()), 10))
	user.Ctime = time.Timestamp{
		Val: &unix,
	}
	user.Roles = []entity.UserRole{
		{
			RoleId: 1,
			UserId: user.UUID,
			Ctime: time.Timestamp{
				Val: &unix,
			},
		},
	}
	user.Data = entity.UserData{
		Name:   name,
		Avatar: "default",
		Ctime: time.Timestamp{
			Val: &unix,
		},
	}
	err := dbTool.Mysql.Create(&user).Error
	return &user, err
}
