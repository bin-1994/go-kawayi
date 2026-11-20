package repository

import (
	"go-biu/app/admin/model"
	"gorm.io/gorm"
	"strings"
	"time"
)

type IBangumiRepository interface {
	Create(*model.BangumiModel) error
	Update(*model.BangumiModel) error
	Delete(id int) error
	FindById(id int) (model.BangumiModel, error)
	FindAll([]string, []any) ([]model.BangumiModel, error)
	//FindByIds(ids []string) ([]model.BangumiModel, error)
	Count([]string, []any) int64
	//FetchBangumiStyle(id int) []int
	//CreateStyle(int, []string) error
	//DeleteStyle(int) error
	//UpdateEpisode(model model.BangumiModel) error
}

type bangumiRepository struct {
	Db *gorm.DB
}

func NewBangumiRepository(Db *gorm.DB) IBangumiRepository {

	return &bangumiRepository{Db: Db}
}

func (repository *bangumiRepository) Create(bangumi *model.BangumiModel) error {
	bangumi.CreateTime = time.Now().Unix()
	bangumi.UpdateTime = 0
	if err := repository.Db.Table("bangumi").Create(&bangumi); err != nil {
		return err.Error
	}
	return nil
}

func (repository *bangumiRepository) Update(bangumi *model.BangumiModel) error {
	if err := repository.Db.Table("bangumi").Where("id = ?", bangumi.Id).Updates(&bangumi); err != nil {
		return err.Error
	}
	return nil
}

func (repository *bangumiRepository) Delete(id int) error {
	if err := repository.Db.Table("bangumi").Where("id = ?", id); err != nil {
		return err.Error
	}
	return nil
}

func (repository *bangumiRepository) FindById(id int) (model.BangumiModel, error) {
	var bangumi model.BangumiModel
	repository.Db.Table("bangumi").Where("id = ?", id).Find(&bangumi)
	return bangumi, nil
}

func (repository *bangumiRepository) FindAll(query []string, args []any) ([]model.BangumiModel, error) {
	var bangumi []model.BangumiModel
	repository.Db.Table("bangumi").Where(strings.Join(query, " AND "), args...).Find(&bangumi)
	return bangumi, nil
}

func (repository *bangumiRepository) Count(query []string, args []any) int64 {
	var total int64
	repository.Db.Table("bangumi").Where(strings.Join(query, " AND "), args...).Count(&total)
	return total
}
