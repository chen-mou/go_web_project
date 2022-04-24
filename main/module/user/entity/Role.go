package entity

import "project/main/tool/time"

type Auth struct {
	Id    int
	Name  string
	Scope string
	Auth  int32
	Ctime time.Timestamp
	Mtime time.Timestamp
}

type Role struct {
	Id     int
	Name   string
	Status string
	Auths  []RoleAuth `gorm:"foreignKey:RoleId"`
	Ctime  time.Timestamp
	Mtime  time.Timestamp
}

type UserRole struct {
	Id     int
	UserId string
	RoleId int
	Expire int32
	Status string
	Role   Role `gorm:"foreignKey:Id"`
	Ctime  time.Timestamp
	Mtime  time.Timestamp
}

type RoleAuth struct {
	Id     int
	RoleId int
	AuthId int
	Auth   Auth `gorm:"foreignKey:Id"`
	Ctime  time.Timestamp
	Mtime  time.Timestamp
}

func (Role) TableName() string {
	return "role_define"
}

func (Auth) TableName() string {
	return "auth_define"
}

func (UserRole) TableName() string {
	return "user_role"
}

func (RoleAuth) TableName() string {
	return "role_auth"
}
