package service

import (
	"fmt"
	"go-biu/app/api/main/model"
	"go-biu/app/api/main/repository"
	"go-biu/utils"
	"time"
)

type IRankService interface {
	Create(string, string, string) error
	FindAll(int64, int64) []model.BangumiListModel
}

type RankService struct {
	Repository repository.IRankRepository
}

func NewRankService(repository repository.IRankRepository) IRankService {
	return &RankService{Repository: repository}
}

func (service *RankService) Create(id string, ep_id string, ip string) error {
	userKey := fmt.Sprintf("play:%v:%v", ip, ep_id)
	if isExist := utils.NewRedis().Exist(userKey); isExist {
		return nil
	}

	service.Repository.IncrBangumi(id)
	service.Repository.IncrEpisode(ep_id)

	hours := int(time.Now().Unix() / 3600)
	key := fmt.Sprintf("rank:hours:%v", hours)

	if err := utils.NewRedis().ZIncrBy(key, 1, id); err != nil {
		return err
	}
	utils.NewRedis().Expire(key, time.Hour*24*14)
	utils.NewRedis().Set(userKey, 1, time.Hour*2)
	return nil
}

func (service *RankService) FindAll(offset, limit int64) []model.BangumiListModel {

	var (
		keys    []string
		weights []float64
	)
	hours := int(time.Now().Unix() / 3600)
	for i := 0; i < 24*7; i++ {
		keys = append(keys, fmt.Sprintf("rank:hours:%v", hours-i))
		weights = append(weights, 1)
	}
	utils.NewRedis().ZUnionStore("rank:bangumi", keys, weights)
	bangumi := utils.NewRedis().ZRevRangeWithScores("rank:bangumi", offset, limit)

	var ids []interface{}
	var scores []float64
	for _, value := range bangumi {
		ids = append(ids, value.Member.(string))
		scores = append(scores, value.Score)
	}
	list, _ := service.Repository.FindByIds(ids)

	for key, _ := range ids {
		list[key].Score = scores[key]
	}

	return list
}
