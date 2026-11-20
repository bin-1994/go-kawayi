package repository

import (
	"go-biu/app/api/main/model"
	"gorm.io/gorm"
)

type ICommentRepository interface {
	FindAll(int64, int, int) []model.CommentModel
	Count(int64) int64
	FindReply(int64, int, int) []model.ReplyModel
	ReplyCount(int64) int64
	Create(*model.BaseComment) error
	Incr(int64) error
	Delete(int64, int64) error
}

type CommentRepository struct {
	Db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) ICommentRepository {
	return &CommentRepository{Db: db}
}

func (repository *CommentRepository) FindAll(id int64, offset, limit int) []model.CommentModel {
	var list []model.CommentModel
	repository.Db.Table("comment").Where("root = 0 AND ep_id = ?", id).Preload("Replys", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Author").Preload("At").Order("cid ASC").Group("root").Limit(20)
	}).Preload("Author").Offset(offset).Limit(limit).Order("cid DESC").Find(&list)
	return list
}

func (repository *CommentRepository) Count(id int64) (total int64) {
	repository.Db.Table("comment").Where("root = 0 AND ep_id = ?", id).Count(&total)
	return
}

func (repository *CommentRepository) FindReply(root int64, offset, limit int) []model.ReplyModel {
	var list []model.ReplyModel
	repository.Db.Table("comment").
		Where("root = ? AND root <> 0", root).
		Preload("Author").
		Preload("At").
		Offset(offset).Limit(limit).
		Order("cid ASC").
		Find(&list)
	return list
}

func (repository *CommentRepository) ReplyCount(root int64) (total int64) {
	repository.Db.Table("comment").Where("root = ? AND root <> 0", root).Count(&total)
	return
}

func (repository *CommentRepository) Create(data *model.BaseComment) error {
	err := repository.Db.Table("comment").Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func (repository *CommentRepository) Incr(root int64) error {
	err := repository.Db.Exec("UPDATE `comment` SET `rcount` = `rcount` + 1 WHERE `cid` = ?", root).Error
	if err != nil {
		return err
	}
	return nil
}

func (repository *CommentRepository) Delete(uid int64, id int64) error {
	err := repository.Db.Table("comment").Delete("uid = ? AND cid = ?", uid, id).Error
	if err != nil {
		return err
	}
	return nil
}

//func (this *CommentRepository) GetList(db *gorm.DB, ep_id int64, offset, limit int) (list []model.Comment, total int64) {
//	db.Table("comment").
//		Where("root = 0 AND ep_id = ?", ep_id).
//		Count(&total).
//		Preload("Replys", func(sql *gorm.DB) *gorm.DB {
//			return sql.Preload("Author").Preload("At").Order("cid ASC").Group("root").Limit(20)
//		}).
//		Preload("Author").
//		Offset(offset).Limit(limit).
//		Order("cid DESC").
//		Find(&list)
//	return
//}

//func (this *CommentRepository) GetReplyList(db *gorm.DB, root int64, offset, limit int) (list []model.Reply, total int64) {
//	db.Table("comment").
//		Where("root = ? AND root <> 0", root).
//		Count(&total).
//		Preload("Author").
//		Preload("At").
//		Offset(offset).Limit(limit).
//		Order("cid ASC").
//		Find(&list)
//	return
//}

//func (this *CommentRepository) Create(db *gorm.DB, params *model.BaseComment) error {
//	if err := db.Create(params).Error; err != nil {
//		return err
//	}
//	return nil
//}

//func (this *CommentRepository) IncReply(db *gorm.DB, root int64) error {
//	if err := db.Exec("UPDATE `comment` SET `rcount` = `rcount` + 1 WHERE `cid` = ?", root).Error; err != nil {
//		return err
//	}
//	return nil
//}

//func (this *CommentRepository) Delete(db *gorm.DB, uid int64, id string) error {
//	result := db.Exec("DELETE FROM comment WHERE uid = ? AND cid = ?", uid, id)
//	if err := result.Error; err != nil {
//		return err
//	}
//	return nil
//}

//func (this *CommentDao) SetFloor(db *gorm.DB, ep_id, root, cid int64) error {
//	if err := db.Exec("UPDATE comment SET floor = (SELECT COUNT(1) FROM comment WHERE ep_id = ? AND root = ?) WHERE cid = ?", ep_id, root, cid).Error; err != nil {
//		return err
//	}
//	return nil
//}
