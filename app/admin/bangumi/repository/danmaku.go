package repository

import (
	"go-biu/app/admin/bangumi/model"
	"gorm.io/gorm"
)

type DanmakuRepository struct {

}

func (this *DanmakuRepository) Find(db *gorm.DB, offset, limit int) (list []model.Danmaku, total int64) {
	db.Table("danmaku").Count(&total).Offset(offset).Limit(limit).Order("id DESC").Find(&list)
	return list, total
}


func (this *DanmakuRepository) Delete(db *gorm.DB, id string) error {
	if err := db.Table("danmaku").Where("id = ?", id).Delete(&model.Danmaku{}).Error; err != nil {
		return err
	}
	return nil
}