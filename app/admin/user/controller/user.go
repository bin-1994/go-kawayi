package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-biu/app/admin/user/model"
	"go-biu/library/database/orm"
	"go-biu/utils"
	"net/http"
)

func FetchAllUser (c *gin.Context) {
	offset, length := utils.GetPage(c.Query("page"), c.Query("pagesize"))

	var (
		userList []model.User
		total int64
	)

	db := orm.Db
	keyword := c.Query("keyword")
	if keyword != "" {
		db = db.Where("phone LIKE ? OR nickname LIKE ?", fmt.Sprintf("%%%s%%", keyword), fmt.Sprintf("%%%s%%", keyword))
	}

	db.Limit(length).Offset(offset).Order("uid DESC").Find(&userList)
	db.Table("user").Count(&total)

	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"message": "success",
		"result": userList,
		"total": total,
	})
}

func FetchUser (c *gin.Context) {
	uid := c.Param("uid")
	var user model.User

	orm.Db.Where("uid = ?", uid).First(&user)

	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"result": user,
	})
}

func EditUser (c *gin.Context) {
	var form model.UserForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code": 1,
			"message": "请输入必填字段",
		})
		return
	}

	if err := orm.Db.Where("uid = ?", form.Uid).Updates(&form).Error; err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code": 1,
			"message": "编辑失败",
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"message": "编辑成功",
	})
}
