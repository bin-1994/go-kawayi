package controller

import (
	"github.com/gin-gonic/gin"
	"go-biu/app/api/member/service"
	"net/http"
	"strings"
)

func Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	//captcha := c.PostForm("captcha")

	token, err := new(service.MemberService).Login(email, password, c.ClientIP())
	if err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code":    1,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "登录成功",
		"token":   token,
	})
}

func Register(c *gin.Context) {
	//code := c.PostForm("regCode")

	email := strings.TrimSpace(c.PostForm("email"))
	password := strings.TrimSpace(c.PostForm("password"))
	nickname := strings.TrimSpace(c.PostForm("nickname"))
	regCode := strings.TrimSpace(c.PostForm("regCode"))

	if err := new(service.MemberService).Register(email, password, nickname, regCode, c.ClientIP()); err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code": 1,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"message": "注册成功",
	})
}

// 检查用户名
func CheckNickname(c *gin.Context) {
	nickname := c.Query("nickname")
	if err := new(service.MemberService).CheckNickname(nickname); err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code": 1,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "昵称可以使用",
	})
}
