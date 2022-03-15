package entity

import "project/main/tool/time"

type User struct {
	Id       int
	UUID     string
	Name     string
	Password string
	Status   string
	Salt     string
	Data     UserData   `gorm:"foreignKey:Id"`
	Roles    []UserRole `gorm:"foreignKey:UserId"`
	Ctime    time.Timestamp
	Mtime    time.Timestamp
}

type UserData struct {
	Id          int
	Name        string
	Avatar      string
	Description string
	Ctime       time.Timestamp
	Mtime       time.Timestamp
}

func (User) TableName() string {
	return "user_base"
}
