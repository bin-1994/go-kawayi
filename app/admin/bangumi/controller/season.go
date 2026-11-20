package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-biu/app/admin/bangumi/model"
	"go-biu/library/database/orm"
	"go-biu/utils"
	"net/http"
)

func FetchAllSeason (c *gin.Context) {
	offset, length := utils.GetPage(c.Query("page"), c.Query("pagesize"))

	var (
		season []model.Season
		total int64
	)
	db := orm.Db

	keyword := c.Query("keyword")
	if keyword != "" {
		db = db.Where("title LIKE ?", fmt.Sprintf("%%%s%%", keyword))
	}
	db.Offset(offset).Limit(length).Order("season_id DESC").Find(&season)
	db.Table("media_season").Count(&total)

	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"result": season,
		"total": total,
	})
}

func FetchSeason (c *gin.Context) {
	season_id := c.Param("season_id")

	var season model.Season
	orm.Db.Where("season_id = ?", season_id).First(&season)

	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"result": season,
	})
}

func AddSeason(c *gin.Context) {
	var form model.Season

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

	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"message": "添加成功",
	})
}

func EditSeason(c *gin.Context) {
	var form model.Season

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code": 1,
			"message": "请输入必填字段",
		})
		return
	}

	if err := orm.Db.Model(&form).Where("season_id = ?", form.SeasonId).Updates(&form).Error; err != nil {
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
