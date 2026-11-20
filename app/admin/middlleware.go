package admin

import (
	"github.com/gin-gonic/gin"
	"go-biu/utils"
	"net/http"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//var logger model.LoggerModel
		//logger.Path = ctx.FullPath()
		//logger.Ip = ctx.ClientIP()
		//logger.Referer = ctx.Request.Referer()
		//logger.UserAgent = ctx.Request.UserAgent()
		//logger.Method = ctx.Request.Method
		//logger.CreateTime = time.Now().Unix()
		//
		//library.NewRedis().HSet("logger", logger.Ip, logger)

		// fmt.Println(logger)
		// fmt.Println(err)

		ctx.Next()
	}
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		Authorization := c.GetHeader("Authorization")

		claims, err := utils.ParseToken("7e7c2a4adf02c63c8fa2fcb95d289482", Authorization)

		//fmt.Println(claims.(jwt.MapClaims)["role"])
		if err != nil {
			c.JSON(http.StatusOK, map[string]interface{}{
				"code":    -404,
				"message": "身份验证失败",
			})
			c.Abort()
			return
		}

		c.Set("uid", claims.Uid)
		c.Next()
	}

}
