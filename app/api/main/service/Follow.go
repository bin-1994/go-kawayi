package service

import (
	"errors"
	"go-biu/app/api/main/model"
	"go-biu/app/api/main/repository"
	"time"
)

type IFollowService interface {
	Create(model.FollowModel) error
	Update(model.FollowModel) error
	Delete(model.FollowModel) error
	FindAll(int64, int, int, int) ([]model.BangumiFollowModel, int64)
}

type FollowService struct {
	Repository repository.IFollowRepository
}

func NewFollowService(repository repository.IFollowRepository) IFollowService {
	return &FollowService{Repository: repository}
}

func (service *FollowService) Create(follow model.FollowModel) error {

	// 先查询是否存在
	if isExist := service.Repository.IsExist(follow.Uid, follow.MediaId); isExist {
		return errors.New("数据已存在")
	}

	follow.CreateTime = time.Now().Unix()
	if err := service.Repository.Create(follow); err != nil {
		return errors.New("追番失败")
	}

	// 追番数+1
	if err := service.Repository.Incr(follow.MediaId); err != nil {
		return errors.New("追番数量增加失败")
	}
	return nil
}

func (service *FollowService) Update(follow model.FollowModel) error {
	if err := service.Repository.Update(follow); err != nil {
		return errors.New("更新进度失败")
	}
	return nil
}

func (service *FollowService) Delete(follow model.FollowModel) error {
	if err := service.Repository.Delete(follow); err != nil {
		return errors.New("取消追番失败")
	}

	if err := service.Repository.Decr(follow.MediaId); err != nil {
		return errors.New("追番数量减少失败")
	}
	return nil
}

func (service *FollowService) FindAll(uid int64, state, offset, limit int) ([]model.BangumiFollowModel, int64) {

	list := service.Repository.FindAll(uid, state, offset, limit)

	total := service.Repository.Count(uid, state)

	return list, total
}

//func (this *FollowService) Create(uid int64, params *model.FollowParams) error {
//	var (
//		fd repository.FollowRepository
//		md repository.MediaRepository
//	)
//	db := orm.Db
//
//	if IsExist := md.Get(db, params.MediaId); IsExist == nil {
//		return errors.New("番剧不存在")
//	}
//
//	if isExist := fd.Get(db, uid, params.SeriesId); isExist != nil {
//		return errors.New("数据已存在")
//	}
//
//	// 没有数据 开始创建
//	//timestamp := time.Now().Unix()
//	follow := &model.MediaFollow{
//		Uid:        uid,
//		SeriesId:   params.SeriesId,
//		MediaId:    params.MediaId,
//		CreateTime: time.Now().Unix(),
//	}
//	db = db.Begin()
//	if err := fd.Create(db, follow); err != nil {
//		db.Rollback()
//		return errors.New("追番失败")
//	}
//
//	//conn := cache.Redis.Get()
//	//defer conn.Close()
//	//
//	//conn.Do("ZINCRBY", fmt.Sprintf("%s:follow_bangumi", uid), timestamp, params.SeriesId)
//	//conn.Do("SADD", fmt.Sprintf("%s:series_fans", params.SeriesId), uid)
//	//
//	//seriesNum, _ := redis.Int64(conn.Do("SCARD", fmt.Sprintf("%s:series_fans", params.SeriesId)))
//
//	if err := md.IncFollow(db, params.SeriesId); err != nil {
//		db.Rollback()
//		return errors.New("追番失败")
//	}
//	db.Commit()
//	return nil
//}

//func (this *FollowService) Delete(uid int64, series_id int) error {
//	fd := repository.FollowRepository{}
//	db := orm.Db
//	db = db.Begin()
//	if err := fd.Delete(db, uid, series_id); err != nil {
//		db.Rollback()
//		return errors.New("操作失败！")
//	}
//
//	//conn := cache.Redis.Get()
//	//defer conn.Close()
//	//
//	//conn.Do("ZREM", fmt.Sprintf("%s:follow_bangumi", uid), series_id)
//	//conn.Do("SREM", fmt.Sprintf("%s:series_fans", series_id), uid)
//	//
//	//seriesNum, _ := redis.Int64(conn.Do("SCARD", fmt.Sprintf("%s:series_fans", series_id)))
//	if err := new(repository.MediaRepository).DecFollow(db, series_id); err != nil {
//		db.Rollback()
//		return errors.New("追番失败")
//	}
//	db.Commit()
//	return nil
//}

//func (this *FollowService) UpdateProgress(uid int64, params *model.FollowParams) error {
//
//	followRep := repository.FollowRepository{}
//	db := orm.Db
//	if isExist := followRep.Get(db, uid, params.SeriesId); isExist == nil {
//		return errors.New("数据不存在")
//	}
//
//	if err := followRep.Update(db, uid, params.SeriesId, params.MediaId, params.EpId, params.Progress, time.Now().Unix(), params.EpIndex); err != nil {
//		return errors.New("保存进度失败")
//	}
//
//	return nil
//}

//func (this *FollowService) Rest(series_id int) error {
//	count := new(repository.FollowRepository).Count(orm.Db, series_id)
//
//	if err := new(repository.MediaRepository).CountFollow(orm.Db, series_id, count); err != nil {
//		return errors.New("统计追番失败")
//	}
//
//	return nil
//}
