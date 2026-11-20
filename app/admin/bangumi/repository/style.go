package repository

import (
	"go-biu/app/admin/bangumi/model"
	"gorm.io/gorm"
)

type StyleRepository struct {

}

func (this *StyleRepository) Find(db *gorm.DB) (list []model.Style) {

	db.Table("style").Order("style_id DESC").Find(&list)

	return list
}

func (this *StyleRepository) FindById(db *gorm.DB, id string) (data *model.Style) {
	db.Table("style").Where("style_id = ?", id).First(&data)
	return nil
}

func (this *StyleRepository) Create(db *gorm.DB, data *model.Style) error {
	if err := db.Table("style").Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func (this *StyleRepository) Update(db *gorm.DB, data *model.Style) error {
	if err := db.Table("style").Where("style_id = ?", data.StyleId).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (this *StyleRepository) Delete(db *gorm.DB, id string) error {
	if err := db.Table("style").Where("style_id = ?", id).Update("state", "0").Error; err != nil {
		return err
	}
	return nil
}
