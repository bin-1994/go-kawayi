package repository

import (
	"go-biu/app/api/main/model"
	"gorm.io/gorm"
)

type DanmakuRepository struct {}

func (this *DanmakuRepository) GetList(db *gorm.DB, ep_id string, offset, limit int) (list []model.DanmakuModel, total int64) {
	db.Table("danmaku").
		Where("ep_id = ?", ep_id).
		Count(&total).
		Order("ctime DESC, stime ASC, id DESC").
		Offset(offset).Limit(limit).
		Find(&list)
	return
}

func (this *DanmakuRepository) Create(db *gorm.DB, data *model.DanmakuModel) error {
	if err := db.Create(data).Error; err != nil {
		return err
	}
	return nil
}