package repository

import (
	"go-biu/app/api/main/model"
	"gorm.io/gorm"
	"strings"
)

type IBangumiRepository interface {
	FindAll([]string, []any, string, int, int) []model.BangumiListModel
	Count([]string, []any) int64
	FindById(int) model.BangumiModel
	EpToId(string) int
	FindSeasonById(int) []model.SeasonModel
}

type BangumiRepository struct {
	Db *gorm.DB
}

func NewBangumiRepository(db *gorm.DB) IBangumiRepository {
	return &BangumiRepository{Db: db}
}

func (repository *BangumiRepository) FindAll(query []string, args []any, order string, offset, limit int) []model.BangumiListModel {
	var list []model.BangumiListModel
	repository.Db.Table("media").Where(strings.Join(query, " AND "), args...).Order(order).Offset(offset).Limit(limit).Find(&list)
	return list
}

func (repository *BangumiRepository) Count(query []string, args []any) int64 {
	var total int64
	repository.Db.Table("media").Where(strings.Join(query, " AND "), args...).Count(&total)
	return total
}

func (repository *BangumiRepository) FindById(id int) model.BangumiModel {
	var bangumi model.BangumiModel
	repository.Db.Table("media").Where("media_id = ? AND state = 1", id).First(&bangumi)
	return bangumi
}

func (repository *BangumiRepository) EpToId(episode string) int {
	var ep model.EpisodeModel
	repository.Db.Table("episode").
		Where("ep_id = ? AND state = 1", episode).
		First(&ep)
	return ep.MediaId
}

func (repository *BangumiRepository) FindSeasonById(id int) []model.SeasonModel {
	var list []model.SeasonModel
	repository.Db.Table("media").Where("season_id = ? AND season_id <> -1", id).Find(&list)
	return list
}
