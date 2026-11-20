package service

import (
	"fmt"
	"go-biu/utils"
	"time"
)

type IRoomService interface {
	Online(string, string) int64
}

type RoomService struct {
}

func NewRoomService() IRoomService {
	return &RoomService{}
}

//func (this *RoomService) Broadcast(comm *websocket.Conn) {
//	for {
//		select {
//			//case :
//
//		}
//	}
//}

func (s *RoomService) Online(id, ip string) int64 {
	now := time.Now().Unix()
	exp := now + 30

	cacheKey := fmt.Sprintf("online:%s", id)

	utils.NewRedis().ZAdd(cacheKey, float64(exp), ip)
	utils.NewRedis().Expire(cacheKey, 3600)
	// 移除过期时间
	utils.NewRedis().ZRemRangeByScore(cacheKey, "0", fmt.Sprintf("%s", now))
	num := utils.NewRedis().ZCard(cacheKey)

	return num
}
