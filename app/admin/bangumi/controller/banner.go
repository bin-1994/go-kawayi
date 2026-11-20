package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-biu/app/admin/bangumi/model"
	"go-biu/library/database/cache"
	"go-biu/library/database/orm"
	"go-biu/utils"
	"net/http"
)

func FetchAllBanner(c *gin.Context) {
	offset, length := utils.GetPage(c.Query("page"), c.Query("pagesize"))

	var (
		banner []model.BannerModel
		total int64
	)
	db := orm.Db

	keyword := c.Query("keyword")
	if keyword != "" {
		db = db.Where("title LIKE ?", fmt.Sprintf("%%%s%%", keyword))
	}
	db.Offset(offset).Limit(length).Order("id DESC").Find(&banner)
	db.Table("banner").Count(&total)

	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"result": banner,
		"total": total,
	})
}

func FetchBanner (c *gin.Context) {
	id := c.Param("id")

	var banner model.BannerModel
	orm.Db.Where("id = ?", id).First(&banner)

	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"result": banner,
	})
}

func AddBanner(c *gin.Context) {
	var form model.BannerModel

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code": 1,
			"message": "请输入必填字段",
		})
		c.Abort()
		return
	}

	if err := orm.Db.Create(&form).Error; err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code": 1,
			"message": "添加失败",
		})
		return
	}

	cache.Delete("banner")

	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"message": "添加成功",
	})
}

func EditBanner(c *gin.Context) {
	var form model.BannerModel

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code": 1,
			"message": "请输入必填字段",
		})
		return
	}

	if err := orm.Db.Model(&form).Updates(&form).Error; err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code": 1,
			"message": "编辑失败",
		})
		return
	}

	cache.Delete("banner")

	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"message": "编辑成功",
	})
}