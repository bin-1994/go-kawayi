package controller

import (
	"github.com/gin-gonic/gin"
	"go-biu/app/admin/bangumi/service"
	"go-biu/utils"
	"net/http"
)

type DanmakuController struct {

}

func (this *DanmakuController) FetchDanmaku(c *gin.Context) {
	var p utils.Paging
	c.ShouldBind(&p)
	offset, limit := p.SetPage(p.Page, p.Pagesize)

	list, total := new(service.DanmakuService).Find(offset, limit)

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"message": "success",
		"result": gin.H{
			"list": list,
			"page": gin.H{
				"num": p.Page,
				"size": p.Pagesize,
				"total": total,
			},
		},
	})
}

func (this *DanmakuController) DelDanmaku(c *gin.Context) {
	id := c.PostForm("id")

	if err := new(service.DanmakuService).Delete(id); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"message": "删除成功",
	})
}