package entity

import "project/main/tool/time"

type Shop struct {
	Id          int
	ShopId      string
	OwnId       string
	Name        string
	Avatar      string
	Description string
	Status      string
	ctime       time.Timestamp
	mtime       time.Timestamp
}

type Commodity struct {
	Id          int
	ShopId      string
	CommodityId string
	name        string
	stock       int
	price       int
	description string
	ctime       time.Timestamp
	mtime       time.Timestamp
}

type Goods struct {
	Id          int
	ShopId      string
	GoodsId     string
	ClassId     string
	Name        string
	Description string
	Price       float64
	Ctime       time.Timestamp
	Mtime       time.Timestamp
}
