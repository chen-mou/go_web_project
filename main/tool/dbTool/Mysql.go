package dbTool

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"sync"
)

var Mysql *gorm.DB = nil

func init() {
	master := "root:CZLczl@20010821@tcp(120.24.214.131:12000)/system"
	slave := "root:CZLczl@20010821@tcp(120.24.214.131:12001)/system"
	once := sync.Once{}
	if Mysql == nil {
		once.Do(func() {
			Mysql, _ = gorm.Open(mysql.New(mysql.Config{DSN: master}), &gorm.Config{})
			Mysql.Use(dbresolver.Register(dbresolver.Config{
				Replicas: []gorm.Dialector{mysql.Open(slave)},
			}))
		})
	}
}
