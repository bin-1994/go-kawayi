package repository

import (
	"go-biu/app/admin/bangumi/model"
	"gorm.io/gorm"
)

type EpisodeRepository struct {

}

func (this *EpisodeRepository) First(db *gorm.DB, id string) (data *model.Episode) {
	db.Table("episode").Where("ep_id = ?", id).First(&data)
	return 
}

func (this *EpisodeRepository) FindAll(db *gorm.DB, offset, limit int) (list []model.Episode, total int64) {
	db.Table("episode").Count(&total).Offset(offset).Limit(limit).Find(&list)
	return list, total
}

func (this *EpisodeRepository) FindById(db *gorm.DB, id string, offset, limit int) (list []model.Episode, total int64) {
	db.Table("episode").Where("media_id = ?", id).Count(&total).Offset(offset).Limit(limit).Order("ep_id DESC").Find(&list)
	return list, total
}

func (this *EpisodeRepository) Create(db *gorm.DB, data *model.EpForm) error {
	if err := db.Table("episode").Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func (this *EpisodeRepository) Update(db *gorm.DB, data model.EpForm) error {
	if err := db.Table("episode").Where("ep_id = ?", data.EpId).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (this *EpisodeRepository) Delete(db *gorm.DB, id string) error {
	if err := db.Table("episode").Where("ep_id = ?", id).Update("state", 0).Error; err != nil {
		return err
	}
	return nil
}

func (this *EpisodeRepository) IsExistEp(db *gorm.DB, media_id, index string) int64 {
	row := db.Table("episode").Where("media_id = ? AND index = ?", media_id ,index).First(&model.Episode{}).RowsAffected
	return row
}
