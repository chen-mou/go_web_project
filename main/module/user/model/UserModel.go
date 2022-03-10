package model

import (
	"golang.org/x/sys/windows"
	"gorm.io/gorm"
	"project/main/module/user/entity"
	"project/main/tool/dbTool"
	"project/main/tool/encryption"
	"project/main/tool/time"
	"strconv"
	local "time"
)

func Register(name string, password string) (*entity.User, string) {
	user := entity.User{}
	err := dbTool.Mysql.Where("name = ?", name).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			unix := local.Now().Unix()
			user.Name = name
			user.Status = "NORMAL"
			user.Salt = encryption.MD5Salt(user.Name, strconv.FormatInt(unix, 10))
			user.Password = encryption.MD5SaltCount(password, user.Salt, 5)
			user.UUID = encryption.MD5Salt(name,
				"machine_name"+strconv.FormatInt(unix, 10)+
					strconv.FormatUint(uint64(windows.GetCurrentProcessId()), 10))
			user.Ctime = time.Timestamp{
				Val: &unix,
			}
			err = dbTool.Mysql.Create(&user).Error
			if err != nil {
				panic(any(err))
			}
			user.Salt = ""
			return &user, "注册成功"
		}
		panic(any(err.Error()))
	}
	return nil, "用户已存在"
}
