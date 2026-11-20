package service

import (
	"errors"
	"go-biu/app/api/main/model"
	"go-biu/app/api/main/repository"
	"strings"
	"time"
)

type ICommentService interface {
	FindAll(int64, int, int) ([]model.CommentModel, int64)
	FindReply(int64, int, int) ([]model.ReplyModel, int64)
	Create(model.BaseComment, int64, string) (model.BaseComment, error)
	Delete(int64, int64) error
}

type CommentService struct {
	Repository repository.ICommentRepository
}

func NewCommentService(repository repository.ICommentRepository) ICommentService {
	return &CommentService{Repository: repository}
}

func (service *CommentService) FindAll(id int64, offset, limit int) ([]model.CommentModel, int64) {
	list := service.Repository.FindAll(id, offset, limit)
	total := service.Repository.Count(id)

	return list, total
}

func (service *CommentService) FindReply(root int64, offset, limit int) ([]model.ReplyModel, int64) {
	list := service.Repository.FindReply(root, offset, limit)
	total := service.Repository.ReplyCount(root)

	return list, total
}

func (service *CommentService) Create(comment model.BaseComment, uid int64, ip string) (model.BaseComment, error) {
	comment.Message = strings.TrimSpace(comment.Message)
	if len(comment.Message) == 0 {
		return comment, errors.New("请输入评论内容")
	}

	data := model.BaseComment{
		Uid:      uid,
		ReplyUid: comment.ReplyUid,
		EpId:     comment.EpId,
		Root:     comment.Root,
		Parent:   comment.Parent,
		Plat:     comment.Plat,
		Message:  comment.Message,
		Ctime:    time.Now().Unix(),
		Ip:       ip,
	}
	if err := service.Repository.Create(&data); err != nil {
		return comment, errors.New("评论失败")
	}

	// 添加回复数量
	if comment.Root != 0 {
		if err := service.Repository.Incr(comment.Root); err != nil {
			return comment, errors.New("评论失败")
		}
	}

	return data, nil
}

func (service *CommentService) Delete(uid int64, id int64) error {
	if err := service.Repository.Delete(uid, id); err != nil {
		return errors.New("删除失败")
	}
	return nil
}

//func (this *CommentService) GetList(ep_id int64, offset, limit int) ([]model.Comment, int64) {
//
//	var (
//		commentRepository repository.CommentRepository
//		list              []model.Comment
//		total             int64
//	)
//	list, total = commentRepository.GetList(orm.Db, ep_id, offset, limit)
//
//	return list, total
//}

//func (this *CommentService) GetReplyList(pid int64, offset, limit int) ([]model.Reply, int64) {
//	list, total := new(repository.CommentRepository).GetReplyList(orm.Db, pid, offset, limit)
//	return list, total
//}

//func (this *CommentService) Create(uid int64, ip string, params *model.BaseComment) (*model.BaseComment, error) {
//	params.Message = strings.TrimSpace(params.Message)
//	if len(params.Message) == 0 {
//		return nil, errors.New("请输入评论内容")
//	}
//
//	comment := model.BaseComment{
//		Uid:      uid,
//		ReplyUid: params.ReplyUid,
//		EpId:     params.EpId,
//		Root:     params.Root,
//		Parent:   params.Parent,
//		Plat:     params.Plat,
//		Message:  params.Message,
//		Ctime:    time.Now().Unix(),
//		Ip:       ip,
//	}
//
//	commentDao := repository.CommentRepository{}
//	db := orm.Db.Begin()
//	if err := commentDao.Create(db, &comment); err != nil {
//		db.Rollback()
//		return nil, errors.New("评论失败")
//	}
//
//	if params.Root != 0 {
//		if err := commentDao.IncReply(db, comment.Root); err != nil {
//			db.Rollback()
//			return nil, errors.New("评论失败")
//		}
//	}
//
//	db.Commit()
//	return &comment, nil
//}

//func (this *CommentService) DelReply(uid int64, id string) error {
//	if err := new(repository.CommentRepository).Delete(orm.Db, uid, id); err != nil {
//		return errors.New("删除失败")
//	}
//	return nil
//}
