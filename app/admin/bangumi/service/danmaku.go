package service

import (
	"errors"
	"go-biu/app/admin/bangumi/model"
	"go-biu/app/admin/bangumi/repository"
	"go-biu/library/database/orm"
)

type DanmakuService struct {

}

func (this *DanmakuService) Find(offset, limit int) (list []model.Danmaku, total int64) {
	list, total = new(repository.DanmakuRepository).Find(orm.Db, offset, limit)
	return list, total
}

func (this *DanmakuService) Delete(id string) error {
	if id == "" {
		return errors.New("参数错误")
	}
	if err := new(repository.DanmakuRepository).Delete(orm.Db, id); err != nil {
		return errors.New("删除失败")
	}
	return nil
}
