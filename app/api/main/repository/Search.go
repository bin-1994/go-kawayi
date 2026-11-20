package repository

import (
	"go-biu/app/api/main/model"
	"gorm.io/gorm"
)

type ISearchRepository interface {
	FindAll(string, bool, int, int) []model.SearchModel
	Count(string) int64
}

type SearchRepository struct {
	Db *gorm.DB
}

func NewSearchRepository(db *gorm.DB) ISearchRepository {
	return &SearchRepository{Db: db}
}

func (repository *SearchRepository) FindAll(keyword string, hasEp bool, offset, limit int) []model.SearchModel {
	var list []model.SearchModel
	db := repository.Db.
		Table("media").
		Where("state = 1 AND MATCH ( `title`, `origin_name`, `keyword` ) AGAINST (? IN BOOLEAN MODE)", "*"+keyword+"*")

	if hasEp {
		db = db.Preload("Episodes", func(db *gorm.DB) *gorm.DB {
			return db.Order("`index` ASC").Where("state = 1")
		})
	}

	db.Offset(offset).Limit(limit).Find(&list)

	return list
}

func (repository *SearchRepository) Count(keyword string) int64 {
	var total int64
	repository.Db.Where("state = 1 AND MATCH ( `title`, `origin_name`, `keyword` ) AGAINST (? IN BOOLEAN MODE)", "*"+keyword+"*").Count(&total)
	return total
}
