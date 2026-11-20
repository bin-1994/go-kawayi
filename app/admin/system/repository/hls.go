package repository

import (
	"fmt"

	"go-biu/app/admin/system/model"
	"gorm.io/gorm"
)

type HlsRepository interface {
	List(*model.HlsModel) []model.HlsModel
	Get(int64) model.HlsModel
	Create(hlsModel *model.HlsModel) (*model.HlsModel, error)
	Update(*model.HlsModel) (*model.HlsModel, error)
	IsExist(hlsModel *model.HlsModel) bool
}

type repository struct {
	Db *gorm.DB
}

func NewHlsRepository(db *gorm.DB) HlsRepository {
	return &repository{
		Db: db,
	}
}

func (r *repository) List(condition *model.HlsModel) (data []model.HlsModel) {
	r.Db.Where(condition).Order("name ASC, id ASC").Find(&data)
	return
}

func (r *repository) Get(id int64) model.HlsModel {
	var data model.HlsModel
	r.Db.Where("id = ?", id).First(&data)
	return data
}

func (r *repository) Create(data *model.HlsModel) (*model.HlsModel, error) {
	if err := r.Db.Create(data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (r *repository) Update(data *model.HlsModel) (*model.HlsModel, error) {
	if err := r.Db.Updates(&data).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}

func (r *repository) IsExist(condition *model.HlsModel) bool {
	row := r.Db.Where(condition).First(&condition).RowsAffected
	if row == 0 {
		return false
	}
	return true
}
