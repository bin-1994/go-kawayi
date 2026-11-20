package repository

import (
	"go-biu/app/api/main/model"
	"gorm.io/gorm"
)

type IEpisodeRepository interface {
	FindAll(int) []model.EpisodeModel
}

type EpisodeRepository struct {
	Db *gorm.DB
}

func NewEpisodeRepository(db *gorm.DB) IEpisodeRepository {
	return &EpisodeRepository{Db: db}
}

func (repository *EpisodeRepository) FindAll(id int) []model.EpisodeModel {
	var list []model.EpisodeModel
	repository.Db.
		Table("episode").
		Where("media_id = ? AND state = 1", id).
		Order("`version` ASC, `index` ASC, ep_id ASC").
		Find(&list)
	return list
}

//func (this *EpisodeRepository) Get(db *gorm.DB, id, is_limit int) (list []model.Episode) {
//	if is_limit > 1 {
//		db = db.Omit("vurl,format_title").Where("version <> 1")
//	}
//	db.Where("media_id = ? AND state = 1", id).Order("version ASC, `index` ASC").Find(&list)
//	return
//}
//
//func (this *EpisodeRepository) EpFindId(db *gorm.DB, ep_id int64) (id int) {
//	db.Table("episode").Select("media_id").Where("ep_id = ?", ep_id).Row().Scan(&id)
//	return
//}
//
//func (this *EpisodeRepository) AddView(db *gorm.DB, ep_id string) error {
//	if err := db.Exec("UPDATE `episode` SET `views` = `views` + 1 WHERE ep_id = ?", ep_id).Error; err != nil {
//		return err
//	}
//	return nil
//}
