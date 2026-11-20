package repository

import (
	"go-biu/app/api/main/model"
	"gorm.io/gorm"
)

type IStyleRepository interface {
	FindByIds(int) []model.StyleModel
	Pluck(int) []int
	FindAll() []model.StyleModel
	FindByBangumiIds(int) []int
}

type StyleRepository struct {
	Db *gorm.DB
}

func NewStyleRepository(db *gorm.DB) IStyleRepository {
	return &StyleRepository{Db: db}
}

// 查询style的所有id
func (repository *StyleRepository) FindByIds(id int) []model.StyleModel {
	var list []model.StyleModel
	repository.Db.Table("style").Find(&list, repository.Pluck(id))
	return list
}

// 查询关联style
func (repository *StyleRepository) Pluck(id int) []int {
	var ids []int
	repository.Db.Table("media_style").Where("media_id = ?", id).Pluck("style_id", &ids)
	return ids
}

func (repository *StyleRepository) FindAll() []model.StyleModel {
	var list []model.StyleModel
	repository.Db.Table("style").Order("style_id ASC").Where("state = 1").Find(&list)
	return list
}

func (repository *StyleRepository) FindByBangumiIds(id int) []int {
	var ids []int
	repository.Db.Table("media_style").Where("style_id = ?", id).Group("media_id").Pluck("media_id", &ids)
	return ids
}
