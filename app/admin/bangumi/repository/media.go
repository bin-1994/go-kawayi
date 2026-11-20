package repository

import (
	"go-biu/app/admin/bangumi/model"
	"gorm.io/gorm"
)

type MediaRepository struct {

}

func (this *MediaRepository) First(db *gorm.DB, id string) (data model.Media) {
	db.Table("media").Where("media_id = ?", id).First(&data)
	return data
}

func (this *MediaRepository) Update(db *gorm.DB, id int, data interface{}) {
	db.Table("media").Where("media_id = ?", id).Updates(data)
}
