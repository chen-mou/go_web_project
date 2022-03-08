package entity

import "project/main/tool/time"

type User struct {
	Id       int
	UUID     string
	Name     string
	Password string
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
