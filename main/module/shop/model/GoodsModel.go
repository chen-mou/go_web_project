package model

import (
	"errors"
	"gorm.io/gorm"
	"project/main/module/shop/entity"
)

func GetById(id string, tx *gorm.DB) (*entity.Goods, error) {
	var goods entity.Goods
	err := tx.Where("goods_id = ?", id).First(&goods).Error
	if err == nil {
		return &goods, nil
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return nil, err
}

func GetByShopId(id string, tx *gorm.DB) ([]entity.Goods, error) {
	var goods []entity.Goods
	err := tx.Where("shop_id = ?", id).Find(&goods).Error
	if err == nil {
		return goods, nil
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return nil, err
}
