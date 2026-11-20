package service

import (
	"go-biu/app/admin/bangumi/model"
	"go-biu/app/admin/bangumi/repository"
	"go-biu/library/database/orm"
)

type CommentService struct {

}

func (this *CommentService) GetList(offset, limit int) (list []model.Comment, total int64) {

	db := orm.Db

	list, total = new(repository.CommentRepository).GetList(db, offset, limit)

	return list, total
}

func (this *CommentService) Delete(id string) error {
	cr := repository.CommentRepository{}
	db := orm.Db
	// 删除评论
	cr.DelComment(db, id)
	cr.DelReply(db, id)

	return nil
}
