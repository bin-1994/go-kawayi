package controller

import (
	"github.com/gin-gonic/gin"
	"go-biu/app/admin/bangumi/service"
	"go-biu/utils"
	"net/http"
)

type CommentController struct {

}

func (this *CommentController) List(c *gin.Context) {
	var p utils.Paging

	c.ShouldBind(&p)

	offset, limit := p.SetPage(p.Page, p.Pagesize)

	list, total := new(service.CommentService).GetList(offset, limit)

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

func (this *CommentController) Get(c *gin.Context) {

}

func (this *CommentController) Create(c *gin.Context) {

}

func (this *CommentController) Delete(c *gin.Context) {
	id := c.PostForm("cid")

	new(service.CommentService).Delete(id)

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"message": "删除成功",
	})
}
