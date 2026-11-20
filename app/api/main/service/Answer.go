package service

import (
	"fmt"
	"go-biu/app/api/main/model"
	"go-biu/app/api/main/repository"
	"go-biu/utils"
	"time"
)

const countdown = 1801

type IAnswerService interface {
	FindAll(string) ([]interface{}, int, time.Time)
}

type AnswerService struct {
	Repository repository.IAnswerRepository
}

func NewAnswerService() IAnswerService {
	return &AnswerService{}
}

func (service *AnswerService) FindAll(ip string) ([]interface{}, int, time.Time) {
	cacheKey := fmt.Sprintf("answer_%s", ip)
	var list []model.AnswerModel

	// 倒计时30分钟
	now := time.Now()

	if exists := utils.NewRedis().Exist(cacheKey); !exists {
		list = service.Repository.FindAll()
		utils.NewRedis().Set(cacheKey, model.AnswerCacheModel{Timestamp: int(now.Unix()), List: list}, countdown)
	}
	var listCache model.AnswerCacheModel
	utils.NewRedis().Get(cacheKey, &listCache)

	result := make([]interface{}, 0)
	for _, val := range listCache.List {
		var option []string
		option = append(option, val.Option1, val.Option2, val.Option3, val.Option4)
		item := make(map[string]interface{})
		item["question"] = val.Question
		item["option"] = option
		result = append(result, item)
	}
	return result, countdown, now
}
