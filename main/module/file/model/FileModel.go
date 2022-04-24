package model

import (
	"errors"
	"gorm.io/gorm"
	"project/main/module/file/entity"
	"project/main/tool/dbTool"
	mt "project/main/tool/time"
	"time"
)

func Create(file entity.File, tx *gorm.DB) (*entity.File, error) {
	now := time.Now().Unix()
	file.Ctime = mt.Timestamp{
		Val: &now,
	}
	err := tx.Create(&file).Error
	return &file, err
}

//
//func Update(file entity.File, tx *gorm.DB)(entity.File, error) {
//
//}

func Update(file entity.File, tx *gorm.DB, fileId string) (*entity.File, error) {
	now := time.Now().Unix()
	file.Mtime = mt.Timestamp{
		Val: &now,
	}
	err := tx.Model(&file).Where("file_id = ?", fileId).Updates(&file).Error
	return &file, err
}

func HasId(id string) (*entity.File, bool, error) {
	var file entity.File
	err := dbTool.Mysql.
		Where("file_id = ? and status != 'DELETED'", id).
		First(&file).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, false, nil
		}
		return nil, false, err
	}
	return &file, true, err
}
