package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"go-biu/library/database/cache"
	"net/http"
)

func FetchSearchList(c *gin.Context) {
	conn := cache.Redis.Get()
	defer  conn.Close()
	res, _ := redis.Values(conn.Do("SMEMBERS", "search:keyword"))

	//data := make(map[string]interface{}, 0)
	var data []interface{}
	for _, value := range res {
		item := make(map[string]interface{}, 0)
		item["keyword"] = string(value.([]byte))
		data = append(data, item)
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"result": data,
	})
}

func DeleteKeyword(c *gin.Context) {

	keyword := c.PostForm("keyword")
	if keyword == "" {
		return
	}

	conn := cache.Redis.Get()
	defer  conn.Close()

	reply, _ := redis.Bool(conn.Do("SREM", "search:keyword", keyword))

	if reply {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code": 0,
			"message": "success",
		})
	}

}

func DelAllKeyword(c *gin.Context) {
	conn := cache.Redis.Get()
	defer  conn.Close()

	reply, _ := redis.Bool(conn.Do("DEL", "search:keyword"))

	if reply {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code": 0,
			"message": "success",
		})
	}
}
