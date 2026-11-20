package repository

import (
	"go-biu/app/api/main/model"
	"gorm.io/gorm"
)

type IFollowRepository interface {
	Create(model.FollowModel) error
	Update(model.FollowModel) error
	Delete(model.FollowModel) error
	FindAll(int64, int, int, int) []model.BangumiFollowModel
	Count(int64, int) int64
	IsExist(int64, int) bool
	Incr(int) error
	Decr(int) error
}

type FollowRepository struct {
	Db *gorm.DB
}

func NewFollowRepository(db *gorm.DB) IFollowRepository {
	return &FollowRepository{Db: db}
}

func (repository *FollowRepository) Create(follow model.FollowModel) error {
	err := repository.Db.Table("media_follow").Create(&follow).Error
	if err != nil {
		return err
	}
	return nil
}

func (repository *FollowRepository) Update(follow model.FollowModel) error {
	err := repository.Db.
		Table("media_follow").
		Where("uid = ? AND media_id = ?", follow.Uid, follow.MediaId).
		Updates(follow).Error
	if err != nil {
		return err
	}
	return nil
}

func (repository *FollowRepository) Delete(follow model.FollowModel) error {
	err := repository.Db.
		Table("media_follow").
		Where("uid = ? AND media_id = ?", follow.Uid, follow.MediaId).
		Delete(model.FollowModel{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (repository *FollowRepository) FindAll(uid int64, state, offset, limit int) []model.BangumiFollowModel {
	var list []model.BangumiFollowModel
	db := repository.Db.Table("media_follow").
		Where("uid = ?", uid)
	if state != 0 {
		db = db.Where("follow_state = ?", state)
	}
	db.Select("media.media_id,media.title,media.cover,media.video_cover,media.is_finish,media.story,media.ep_id,media_follow.*").
		Joins("INNER JOIN media ON media_follow.media_id = media.media_id").
		Offset(offset).Limit(limit).
		Order("GREATEST(media.update_time, media_follow.create_time) DESC").
		Find(&list)

	return list
}

func (repository *FollowRepository) Count(uid int64, state int) int64 {
	var total int64
	db := repository.Db.Table("media_follow").Where("uid = ?", uid)
	if state != 0 {
		db.Where("follow_state = ?", state)
	}
	db.Count(&total)
	return total
}

func (repository *FollowRepository) IsExist(uid int64, media_id int) bool {
	var follow model.FollowModel
	result := repository.Db.Table("media_follow").
		Where("uid = ? AND media_id = ?", uid, media_id).
		First(&follow)
	if result.Error == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

func (repository *FollowRepository) Incr(id int) error {
	query := "UPDATE media SET follows = follows + 1 WHERE media_id = ?"
	repository.Db.Exec(query, id)
	return nil
}

func (repository *FollowRepository) Decr(id int) error {
	query := "UPDATE media SET follows = follows - 1 WHERE media_id = ? AND follows > 0"
	repository.Db.Exec(query, id)
	return nil
}

//func (this *FollowRepository) Get(db *gorm.DB, uid int64, series_id int) *model.MediaFollow {
//	ret := &model.MediaFollow{}
//	if err := db.Where("uid = ? AND series_id = ?", uid, series_id).First(ret).Error; err != nil {
//		return nil
//	}
//	return ret
//}

//func (this *FollowRepository) Create(db *gorm.DB, data *model.MediaFollow) (err error) {
//	err = db.Create(data).Error
//	return
//}

//func (this *FollowRepository) Update(db *gorm.DB, uid int64, series_id, media_id, ep_id, progress int, last_time int64, ep_index string) error {
//	if err := db.Table("media_follow").Where("uid = ? AND series_id = ?", uid, series_id).Updates(map[string]interface{}{
//		"media_id": media_id,
//		"ep_id": ep_id,
//		"ep_index": ep_index,
//		"progress": progress,
//		"last_time": last_time,
//	}).Error; err != nil {
//		return err
//	}
//	return nil
//}

//func (this *FollowRepository) Delete(db *gorm.DB, uid int64, series_id int) (err error) {
//	err = db.Where("uid = ? AND series_id = ?", uid, series_id).Delete(&model.MediaFollow{}).Error
//	return
//}

//func (this *FollowRepository) Count(db *gorm.DB, id int) int64 {
//	var count int64
//	db.Table("media_follow").Where("series_id = ?", id).Count(&count)
//	return count
//}
