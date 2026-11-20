package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-biu/app/admin/bangumi/model"
	"go-biu/library/database/orm"
	"go-biu/utils"
	"net/http"
)

func FetchAllAnswer(c *gin.Context) {
	offset, length := utils.GetPage(c.Query("page"), c.Query("pagesize"))

	var (
		answer []model.Answer
		total int64
	)

	db := orm.Db
	keyword := c.Query("keyword")
	if keyword != "" {
		db = db.Where("question LIKE ?", fmt.Sprintf("%%%s%%", keyword))
	}

	db.Limit(length).Offset(offset).Order("qid DESC").Find(&answer)

	db.Table("answer").Count(&total)

	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"result": answer,
		"total": total,
	})
}

func FetchAnswer(c *gin.Context) {
	qid := c.Param("qid")

	var answer model.Answer
	orm.Db.Where("qid = ?", qid).First(&answer)

	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"result": answer,
	})
}

func AddAnswer(c *gin.Context) {
	var form model.AnswerForm

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

func EditAnswer(c *gin.Context) {
	var form model.AnswerForm

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

	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"message": "编辑成功",
	})
}
