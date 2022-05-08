package entity

import (
	"project/main/tool/dbTool"
	"project/main/tool/time"
)

type Order struct {
	Id         int
	UserId     string
	OrderId    string
	BasePrice  float64
	Discount   float64
	FinalPrice float64
	Ctime      time.Timestamp
	mtime      time.Timestamp
}

type OrderTableName struct {
	Id             int
	StartTime      time.Time
	EndTime        time.Time
	TableName      string
	GoodsTableName string
}

type OrderGoods struct {
	Id         int
	OrderId    string
	GoodsId    string
	Discount   float64
	BasePrice  float64
	FinalPrice float64
	Num        int
	Ctime      time.Timestamp
	Mtime      time.Timestamp
}

func (o Order) TableName() string {
	if o.Ctime.Val == nil {
		var tableNames []OrderTableName
		dbTool.Mysql.Find(&tableNames)
		name := tableNames[0].TableName
		for _, names := range tableNames {
			if names.TableName == tableNames[0].TableName {
				continue
			}
			name += ", " + names.TableName
		}
		return name
	}
	var tableName OrderTableName
	dbTool.Mysql.
		Where("start_time < ? and end_time > ?", o.Ctime.Val, o.Ctime.Val).
		First(&tableName)
	return tableName.TableName
}

func (o OrderGoods) TableName() string {
	if o.Ctime.Val == nil {
		var tableNames []OrderTableName
		dbTool.Mysql.Find(&tableNames)
		name := tableNames[0].GoodsTableName
		for _, names := range tableNames {
			if names.GoodsTableName == tableNames[0].GoodsTableName {
				continue
			}
			name += ", " + names.GoodsTableName
		}
		return name
	}
	var tableName OrderTableName
	dbTool.Mysql.
		Where("start_time < ? and end_time > ?", o.Ctime.Val, o.Ctime.Val).
		First(&tableName)
	return tableName.GoodsTableName
}
