package repository

import (
	"go-biu/app/api/main/model"
	"gorm.io/gorm"
)

type ReviewRepository struct {

}

func (this *ReviewRepository) Get(db *gorm.DB, uid int64, id int) (review *model.Review) {
	if err := db.Table("review").Where("uid = ? AND media_id = ?", uid, id).First(&review).Error; err != nil {
		return nil
	}
	return
}

func (this *ReviewRepository) GetList(db *gorm.DB, id int) (list []model.Review) {
	db.Table("review").Where("media_id = ?", id).Offset(0).Limit(3).Find(&list)
	return
}

func (this *ReviewRepository) Create(db *gorm.DB, params *model.ReviewForm) error {
	if err := db.Table("review").Create(&params).Error; err != nil {
		return err
	}
	return nil
}

func (this *ReviewRepository) Update(db *gorm.DB, uid int64, params *model.ReviewForm) error {
	if err := db.Table("review").Where("review_id = ? AND uid = ?", params.ReviewId, uid).Updates(&params).Error; err != nil {
		return err
	}
	return nil
}

func (this *ReviewRepository) Delete(db *gorm.DB, uid int64, media_id int, review_id int64) error {
	if err := db.Table("review").Where("uid = ? AND media_id = ? AND review_id = ?", uid, media_id, review_id).Error; err != nil {
		return err
	}
	return nil
}

func (this *ReviewRepository) Count(db *gorm.DB, media_id int) (int64, int64) {
	result := struct {
		Total int64
		Score int64
	}{}
	db.Table("review").
		Select("count(1) AS total, sum(score) AS score").
		Where("media_id = ?", media_id).
		Scan(&result)
	return result.Total, result.Score
}