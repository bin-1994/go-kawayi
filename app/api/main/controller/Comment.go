package controller

import (
	"github.com/gin-gonic/gin"
	"go-biu/app/api/main/model"
	"go-biu/app/api/main/service"
	"go-biu/utils"
	"strconv"
)

type CommentController struct {
	Service service.ICommentService
}

func NewCommentController(service service.ICommentService) *CommentController {
	return &CommentController{Service: service}
}

func (controller *CommentController) FindAll(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Query("ep_id"), 10, 64)
	resp := utils.NewResponse(ctx)

	page := utils.Paging{Page: ctx.Query("page"), Size: ctx.Query("size")}
	offset, limit := page.SetPage()

	list, total := controller.Service.FindAll(id, offset, limit)

	resp.List(list, offset, limit, total)
}

func (controller *CommentController) FindReply(ctx *gin.Context) {
	//id, _ := strconv.ParseInt(ctx.DefaultQuery("ep_id", "0"), 10, 64)
	root, _ := strconv.ParseInt(ctx.DefaultQuery("root", "0"), 10, 64)

	resp := utils.NewResponse(ctx)

	page := utils.Paging{Page: ctx.Query("page"), Size: ctx.Query("size")}
	offset, limit := page.SetPage()

	list, total := controller.Service.FindReply(root, offset, limit)

	resp.List(list, offset, limit, total)
}

func (controller *CommentController) Create(ctx *gin.Context) {
	var comment model.BaseComment
	resp := utils.NewResponse(ctx)

	uid, _ := ctx.Get("uid")

	if err := ctx.ShouldBind(&comment); err != nil {
		resp.Error("评论失败")
		return
	}
	comment, err := controller.Service.Create(comment, uid.(int64), ctx.ClientIP())
	if err != nil {
		resp.Error("评论失败")
		return
	}

	resp.Success(comment)
}

func (controller *CommentController) Delete(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.PostForm("id"), 10, 64)

	uid, _ := ctx.Get("uid")

	resp := utils.NewResponse(ctx)

	if err := controller.Service.Delete(uid.(int64), id); err != nil {
		resp.Error("删除失败")
		return
	}

	resp.Success("删除成功")
}

//func FetchComment (c *gin.Context) {
//	id := c.DefaultQuery("ep_id", "0")
//
//	p := utils.Paging{}
//	if err := c.ShouldBind(&p); err != nil {
//		c.JSON(http.StatusOK, gin.H{
//			"code": 1,
//			"message": "参数错误",
//		})
//		return
//	}
//	//page := c.DefaultQuery("page", "1")
//	//pagesize := c.DefaultQuery("pagesize", "20")
//	start, end := p.SetPage(p.Page, p.Pagesize)
//	ep_id, _ := strconv.ParseInt(id, 10, 64)
//
//	list, total := new(service.CommentService).GetList(ep_id, start, end)
//
//	c.JSON(http.StatusOK, map[string]interface{}{
//		"code": 0,
//		"message": "success",
//		"result": map[string]interface{}{
//			"list": list,
//			"page": map[string]interface{}{
//				"num": p.Page,
//				"size": p.Pagesize,
//				"total": total,
//			},
//		},
//	})
//}

//func FetchReply (c *gin.Context) {
//	//ep_id := c.DefaultQuery("ep_id", "0")
//	root := c.DefaultQuery("root", "0")
//	//page := c.DefaultQuery("page", "1")
//	//pagesize := c.DefaultQuery("pagesize", "20")
//	pid, _ := strconv.ParseInt(root, 10, 64)
//	p := utils.Paging{}
//	if err := c.ShouldBind(&p); err != nil {
//		c.JSON(http.StatusOK, gin.H{
//			"code": 1,
//			"message": "参数错误",
//		})
//		return
//	}
//	start, end := p.SetPage(p.Page, p.Pagesize)
//
//	//var (
//	//	reply []model.Reply
//	//	count int64
//	//)
//	//result := orm.Db.Preload("Author").Where("ep_id = ? AND root = ? AND state = 1", ep_id, root).Order("id ASC").Offset(start).Limit(end).Find(&reply).Count(&count)
//	//if result.Error != nil {
//	//	c.JSON(http.StatusOK, map[string]interface{}{
//	//		"code": -404,
//	//		"message": "地球上找不到该内容哦 (╥﹏╥)",
//	//	})
//	//	return
//	//}
//
//	list, total := new(service.CommentService).GetReplyList(pid, start, end)
//
//	c.JSON(http.StatusOK, map[string]interface{}{
//		"code": 0,
//		"message": "success",
//		"result": map[string]interface{}{
//			"list": list,
//			"page": map[string]interface{}{
//				"num": p.Page,
//				"size": p.Pagesize,
//				"total": total,
//			},
//		},
//	})
//}

//func AddReply (c *gin.Context) {
//	uid, _ := c.Get("uid")
//	ip := c.ClientIP()
//
//	var commentForm model.BaseComment
//	if err := c.ShouldBind(&commentForm); err != nil {
//		c.JSON(http.StatusOK, map[string]interface{}{
//			"code": 1,
//			"message": "数据错误",
//		})
//		return
//	}
//
//	comment, err := new(service.CommentService).Create(int64(uid.(float64)), ip, &commentForm)
//	if err != nil {
//		c.JSON(http.StatusOK, map[string]interface{}{
//			"code": 1,
//			"message": err.Error(),
//		})
//		return
//	}
//
//	c.JSON(http.StatusOK, map[string]interface{}{
//		"code": 0,
//		"message": "评论成功",
//		"result": comment,
//	})
//}

//func DelReply(c *gin.Context) {
//	uid, _ := c.Get("uid")
//
//	id := c.PostForm("id")
//
//	if err := new(service.CommentService).DelReply(int64(uid.(float64)), id); err != nil {
//		c.JSON(http.StatusOK, gin.H{
//			"code": 1,
//			"message": err.Error(),
//		})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"code": 0,
//		"message": "删除成功",
//	})
//}
