package controller

import (
	"github.com/gin-gonic/gin"
	"go-biu/app/api/main/model"
	"go-biu/app/api/main/service"
	"go-biu/utils"
	"time"
)

//var FeedbackController = newFeedbackController()

//func newFeedbackController() *feedbackController {
//	return &feedbackController{}
//}

type FeedbackController struct {
	Service service.IFeedbackService
}

func NewFeedbackController(service service.IFeedbackService) *FeedbackController {
	return &FeedbackController{Service: service}
}

func (controller *FeedbackController) Create(ctx *gin.Context) {
	var feedback model.FeedbackModel

	resp := utils.NewResponse(ctx)

	if err := ctx.ShouldBind(&feedback); err != nil {
		resp.Error("数据错误")
		return
	}

	feedback.Ip = ctx.ClientIP()
	feedback.UserAgent = ctx.Request.UserAgent()
	feedback.CreateTime = time.Now().Unix()

	if err := controller.Service.Create(feedback); err != nil {
		resp.Error(err.Error())
		return
	}
	
	resp.Success("感谢你提交的反馈")
}

//func (controller *FeedbackController) PostFeedback(c *gin.Context) {
//	media_id := c.PostForm("media_id")
//	ep_id := c.PostForm("ep_id")
//	reason := c.PostForm("reason")
//	message := c.PostForm("message")
//	contact := c.PostForm("contact")
//	ip := c.ClientIP()
//	ua := c.GetHeader("User-Agent")
//
//	feedbackDao := repository.FeedbackRepository{}
//
//	if exists := feedbackDao.IsExits(media_id); exists {
//		c.JSON(http.StatusOK, map[string]interface{}{
//			"code":    1,
//			"message": "问题已存在，请不要重复提交。",
//		})
//		return
//	}
//
//	result := feedbackDao.Add(media_id, ep_id, reason, message, contact, ua, ip, time.Now().Unix())
//
//	if result != nil {
//		c.JSON(http.StatusOK, map[string]interface{}{
//			"code":    1,
//			"message": "提交反馈失败",
//		})
//		return
//	}
//
//	c.JSON(http.StatusOK, map[string]interface{}{
//		"code":    0,
//		"message": "提交反馈成功",
//	})
//}
//
//func (this *feedbackController) VideoFeedback(c *gin.Context) {
//	var p model.FeedbackVideo
//	c.ShouldBind(&p)
//
//	p.Ip = c.ClientIP()
//	p.Ua = c.GetHeader("User-Agent")
//	p.Ctime = time.Now().Unix()
//
//	if err := new(service.FeedbackService).Craete(p); err != nil {
//		c.JSON(http.StatusOK, map[string]interface{}{
//			"code":    1,
//			"message": "提交反馈失败",
//		})
//		return
//	}
//
//	c.JSON(http.StatusOK, map[string]interface{}{
//		"code":    0,
//		"message": "提交反馈成功",
//	})
//}
