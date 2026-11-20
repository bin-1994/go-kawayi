package service

import (
	"fmt"
	"go-biu/app/api/main/model"
	"go-biu/app/api/main/repository"
	"go-biu/utils"
	"time"
)

type ISearchService interface {
	Create(string, float64)
	FindHot(int64, int64) []any
	FindAll(string, bool, int, int) ([]model.SearchModel, int64)
}

type SearchService struct {
	Repository repository.ISearchRepository
}

func NewSearchService(repository repository.ISearchRepository) ISearchService {
	return &SearchService{Repository: repository}
}

// 创建热搜词
func (service *SearchService) Create(keyword string, score float64) {
	timestamp := utils.CurrentTimestamp()
	key := fmt.Sprintf("search:keyword:%v", timestamp)
	utils.NewRedis().ZIncrBy(key, score, keyword)
	utils.NewRedis().Expire(key, time.Hour*24*7)
}

// 查询热搜词
func (service *SearchService) FindHot(offset, limit int64) []any {
	timestamp := utils.CurrentTimestamp()

	keys := []string{fmt.Sprintf("search:keyword:%v", timestamp-(86400*2)), fmt.Sprintf("search:keyword:%v", timestamp-86400), fmt.Sprintf("search:keyword:%v", timestamp)}

	utils.NewRedis().ZUnionStore("search:hot", keys, []float64{1, 1, 1})
	keywords := utils.NewRedis().ZRevRangeWithScores("search:hot", offset, limit)
	var list []any
	for _, value := range keywords {
		item := make(map[string]interface{}, 0)
		item["keyword"] = value.Member
		item["score"] = value.Score
		list = append(list, item)
	}
	if list == nil {
		return []any{}
	}
	return list
}

// 搜索列表
func (service *SearchService) FindAll(keyword string, hasEp bool, offset, limit int) ([]model.SearchModel, int64) {
	list := service.Repository.FindAll(keyword, hasEp, offset, limit)
	var total int64
	if hasEp {
		total = service.Repository.Count(keyword)
	}

	return list, total
}
