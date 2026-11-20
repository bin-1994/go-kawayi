package controller

import (
	"github.com/gin-gonic/gin"
	"go-biu/app/api/main/service"
	"go-biu/utils"
)

//func OnlineRoom(c *gin.Context) {
//	id := c.Param("id")
//	ip := c.ClientIP()
//
//	now := time.Now().Unix()
//	exp := now + 30
//	conn := cache.Redis.Get()
//	defer conn.Close()
//	cacheKey := fmt.Sprintf("online:%s", id)
//	conn.Do("ZADD", cacheKey, exp, ip)
//	// 保留一小时
//	conn.Do("EXPIRE", cacheKey, 3600)
//	conn.Do("ZREMRANGEBYSCORE", cacheKey, 0, now)
//	num, err := redis.Int64(conn.Do("ZCARD", cacheKey))
//
//	if err != nil {
//		panic(err)
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"code": 0,
//		"message": "获取人数成功",
//		"result": num,
//	})
//}

type RoomController struct {
	Service service.IRoomService
}

func NewRoomController(service service.IRoomService) *RoomController {
	return &RoomController{Service: service}
}

func (controller *RoomController) Online(ctx *gin.Context) {
	id := ctx.Param("id")
	ip := ctx.ClientIP()

	num := controller.Service.Online(id, ip)

	resp := utils.NewResponse(ctx)

	resp.Success(num)
}
