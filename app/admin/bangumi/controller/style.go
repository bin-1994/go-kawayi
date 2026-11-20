package controller

import (
	"github.com/gin-gonic/gin"
	"go-biu/app/admin/bangumi/model"
	"go-biu/app/admin/bangumi/service"
	"net/http"
)

type StyleController struct {

}

func (this *StyleController) FetchAllStyle (c *gin.Context) {

	list := new(service.StyleService).Find()

	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"result": gin.H{
			"list": list,
		},
	})
}

func (this *StyleController) Create(c *gin.Context) {
	var p model.Style
	if err := c.ShouldBind(&p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"message": "操作失败",
		})
		return
	}

	if err := new(service.StyleService).Create(p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"message": "创建成功",
	})
}

func (this *StyleController) Update(c *gin.Context) {
	var p model.Style
	if err := c.ShouldBind(&p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"message": "操作失败",
		})
		return
	}

	if err := new(service.StyleService).Update(p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"message": "修改成功",
	})
}
