package controller

import (
	"github.com/gin-gonic/gin"
	"go-biu/app/admin/user/model"
	"go-biu/library/database/orm"
	"go-biu/utils"
	"net/http"
	"time"
)

func Login (c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	var user model.User

	if result := orm.Db.Where("username = ? AND username <> '' AND role = 1 AND state = 1", username).First(&user); result.RowsAffected == 0 {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code": 1,
			"message": "用户不存在",
		})
		return
	}

	if utils.UserMd5(password) != user.Password {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code": 1,
			"message": "密码不正确",
		})
		return
	}

	if 1 != user.Role {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code": 1,
			"message": "该用户没权限",
		})
		return
	}

	if 0 == user.State {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code": 1,
			"message": "该用户已被封印",
		})
		return
	}

	claims := make(map[string]interface{})
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(24) * 30).Unix()
	claims["uid"] = user.Uid
	claims["role"] = user.Role

	token := utils.CreateToken("0b080119cbf1138edfa9132471e1a661", claims)

	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"message": "success",
		"token": token,
	})
}
