package repository

import (
	"go-biu/app/admin/bangumi/model"
	"gorm.io/gorm"
)

type CommentRepository struct {

}

func (this *CommentRepository) GetList(db *gorm.DB, offset, limit int) (list []model.Comment, total int64) {
	db.Table("comment").
		Count(&total).
		Offset(offset).Limit(limit).Order("cid DESC").Find(&list)
	return list, total
}

func (this *CommentRepository) DelComment(db *gorm.DB, id string) {
	db.Table("comment").Where("cid = ?", id).Delete(&model.Comment{})
}

func (this *CommentRepository) DelReply(db *gorm.DB, id string) {
	db.Table("comment").Where("root = ?", id).Delete(&model.Comment{})
}