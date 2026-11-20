package repository

import (
	"go-biu/app/admin/model"
	"gorm.io/gorm"
)

type IStyleRepository interface {
	Create(*model.StyleModel) error
	Update(*model.StyleModel) error
	FindAll() []model.StyleModel
}

type styleRepository struct {
	Db *gorm.DB
}

func NewStyleRepository(Db *gorm.DB) IStyleRepository {
	return &styleRepository{Db: Db}
}

func (repository *styleRepository) Create(style *model.StyleModel) error {
	if err := repository.Db.Table("style").Create(style); err != nil {
		return err.Error
	}
	return nil
}

func (repository *styleRepository) Update(style *model.StyleModel) error {
	if err := repository.Db.Table("style").Where("id = ?", style.Id).Updates(style); err != nil {
		return err.Error
	}
	return nil
}

func (repository *styleRepository) FindAll() []model.StyleModel {
	var list []model.StyleModel
	repository.Db.Table("style").Order("letter ASC").Find(&list)
	return list
}
