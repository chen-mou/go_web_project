package entity

import "project/main/tool/time"

type Auth struct {
	Id    int
	Name  string
	Scope string
	Auth  string
	Ctime time.Timestamp
	Mtime time.Timestamp
}

type Role struct {
	Id     int
	Name   string
	Status string
	Ctime  time.Timestamp
	Mtime  time.Timestamp
}

type UserRole struct {
	Id     int
	UserId string
	RoleId int
	Ctime  time.Timestamp
	Mtime  time.Timestamp
}

type RoleAuth struct {
	Id     int
	RoleId int
	AuthId int
	Ctime  time.Timestamp
	Mtime  time.Timestamp
}
