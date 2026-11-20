package service

import (
	"encoding/json"
	"errors"
	"github.com/gomodule/redigo/redis"
	"go-biu/app/admin/bangumi/model"
	"go-biu/library/database/cache"
)

type FeedbackService struct {
	
}

func (this *FeedbackService) FindList() interface{} {
	conn := cache.Redis.Get()
	defer  conn.Close()

	res, _ := redis.Values(conn.Do("HGETALL", "feedback"))

	var data []interface{}
	var feedback model.FeedbackVideo
	for key, value := range res {
		if key % 2 == 1 {
			//item := make(map[string]interface{}, 0)
			//item["keyword"] = string()
			json.Unmarshal(value.([]byte), &feedback)
			data = append(data, feedback)
		}
	}

	return data
}

func (this *FeedbackService) Delete(field string) error {
	conn := cache.Redis.Get()
	defer  conn.Close()

	if _, err := conn.Do("HDEL", "feedback", field); err != nil {
		return errors.New("删除失败")
	}

	return nil
}

func (this *FeedbackService) DelAll() error {
	conn := cache.Redis.Get()
	defer  conn.Close()

	if _, err := redis.Bool(conn.Do("DEL", "feedback")); err != nil {
		return errors.New("删除失败")
	}
	return nil
}