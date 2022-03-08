package dbTool

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"sync"
)

var Mysql *gorm.DB = nil

func Init() {
	once := sync.Once{}
	if Mysql == nil {
		once.Do(func() {
			Mysql, _ = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
		})
	}
}
