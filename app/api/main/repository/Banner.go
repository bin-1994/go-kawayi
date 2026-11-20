package repository

import (
	"go-biu/app/admin/bangumi/model"
	"gorm.io/gorm"
)

type IBannerRepository interface {
	FindAll(int, int) []model.BannerModel
}

type BannerRepository struct {
	Db *gorm.DB
}

func NewBannerRepository(db *gorm.DB) IBannerRepository {
	return &BannerRepository{Db: db}
}

func (repository *BannerRepository) FindAll(offset, limit int) []model.BannerModel {
	var list []model.BannerModel
	repository.Db.Table("banner").
		Where("state = 1").
		Offset(offset).
		Limit(limit).
		Find(&list)
	return list
}
