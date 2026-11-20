package repository

import (
	"go-biu/app/api/main/model"
	"gorm.io/gorm"
)

type IAnswerRepository interface {
	FindAll() []model.AnswerModel
}

type AnswerRepository struct {
	Db *gorm.DB
}

func NewAnswerRepository(db *gorm.DB) IAnswerRepository {
	return &AnswerRepository{Db: db}
}

func (repository *AnswerRepository) FindAll() []model.AnswerModel {
	var list []model.AnswerModel
	repository.Db.Table("answer").Order("RAND()").Limit(50).Find(&list)
	return list
}
