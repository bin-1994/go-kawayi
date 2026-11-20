package api

import (
	"go-biu/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

const secret = "21d6e79b3b2a0f9ad0373441e4f4a03e"

const refresh_secret = "68d33f3c48149a45afd5f76d22c3875c"
const access_secret = "58df9b68b41e798923fe3ab883860b6f"

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, _ := c.Cookie("access_token")

		claims, err := utils.ParseToken(access_secret, tokenString)

		if err != nil {
			c.JSON(http.StatusOK, map[string]interface{}{
				"code":    401,
				"message": "Token验证失败",
			})
			c.Abort()
			return
		}
		c.Set("uid", claims.Uid)
		c.Next()
	}
}

// var ipdb *ip2location.DB
//
//	func init() {
//		var err error
//		ipdb, err = ip2location.OpenDB("./plugins/ip2location/IP2LOCATION-LITE-DB3.IPV6.BIN")
//		if err != nil {
//			return
//		}
//	}
func super() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, _ := c.Cookie("token")
		claims, err := utils.ParseToken(secret, tokenString)

		if err != nil {
			c.Set("uid", claims.Uid)
		}
		c.Next()
	}
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {

		//if _, err := c.Cookie("pass"); err != nil  {
		//	results, err := ipdb.Get_all(c.ClientIP())
		//	if err != nil {
		//		c.JSON(http.StatusOK, map[string]interface{}{
		//			"code": 403,
		//			"message": "没有权限",
		//		})
		//		c.Abort()
		//		return
		//	}
		//	city := strings.ToLower(results.City)
		//	if "beijing" == city || "shanghai" == city || "zhuhai" == city {
		//		c.JSON(http.StatusOK, map[string]interface{}{
		//			"code": 403,
		//			"message": "没有权限",
		//		})
		//		c.Abort()
		//		return
		//	}
		//}

		origin := c.Request.Header.Get("Origin")
		if origin == "http://127.0.0.1:8080" || origin == "http://185.188.6.70:1235" || origin == "http://127.0.0.1:6060" || origin == "http://192.168.0.100:6060" || origin == "http://dev.kawayi.cc:8080" {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			c.Header("Access-Control-Allow-Headers", "X-Requested-With,content-type")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, Authorization")
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//放行所有OPTIONS方法
		if c.Request.Method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		c.Next()
	}
}
