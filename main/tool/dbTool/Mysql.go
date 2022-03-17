package dbTool

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"project/main/module/user/entity"
	myTime "project/main/tool/time"
	"sync"
	"time"
)

var Mysql *gorm.DB = nil

func initRole() {
	val := time.Now().Unix()
	now := myTime.Timestamp{
		Val: &val,
	}
	authRole := []entity.RoleAuth{
		{
			RoleId: 4,
			AuthId: 12,
			Ctime:  now,
		},
		{
			RoleId: 4,
			AuthId: 14,
			Ctime:  now,
		},
		{
			RoleId: 4,
			AuthId: 8,
			Ctime:  now,
		},
		{
			RoleId: 4,
			AuthId: 11,
			Ctime:  now,
		},
		{
			RoleId: 4,
			AuthId: 15,
			Ctime:  now,
		},
	}
	Mysql.Create(&authRole)
}

func init() {
	master := "root:CZLczl@20010821@tcp(120.24.214.131:12000)/system"
	slave := "root:CZLczl@20010821@tcp(120.24.214.131:12001)/system"
	once := sync.Once{}
	if Mysql == nil {
		once.Do(func() {
			Mysql, _ = gorm.Open(mysql.New(mysql.Config{DSN: master}), &gorm.Config{})
			db, _ := Mysql.DB()
			db.SetConnMaxLifetime(20 * time.Minute)
			db.SetMaxIdleConns(10)
			Mysql.Use(dbresolver.Register(dbresolver.Config{
				Replicas: []gorm.Dialector{mysql.Open(slave)},
			}))
		})
	}
	//initRole()
}
