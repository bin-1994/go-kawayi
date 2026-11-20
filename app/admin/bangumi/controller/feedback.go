package controller

import (
	"github.com/gin-gonic/gin"
	"go-biu/app/admin/bangumi/service"
	"net/http"
)

//func FetchAllFeedback(c *gin.Context) {
//	offset, length := utils.GetPage(c.Query("page"), c.Query("pagesize"))
//
//	var (
//		feedback []model.FeedbackModel
//		total int64
//	)
//
//	orm.Db.Offset(offset).Order("id DESC").Limit(length).Find(&feedback)
//	orm.Db.Table("feedback").Count(&total)
//
//	c.JSON(http.StatusOK, map[string]interface{}{
//		"code": 0,
//		"result": feedback,
//		"total": total,
//	})
//}
//
//func DelFeedback(c *gin.Context) {
//	id := c.PostForm("id")
//
//	if id == "" {
//		c.JSON(http.StatusOK, map[string]interface{}{
//			"code": 1,
//			"message": "删除失败",
//		})
//		return
//	}
//
//	if result := orm.Db.Exec("DELETE FROM feedback WHERE id = ?", id); result.Error != nil {
//		c.JSON(http.StatusOK, map[string]interface{}{
//			"code": 1,
//			"message": "删除失败",
//		})
//		return
//	}
//
//	c.JSON(http.StatusOK, map[string]interface{}{
//		"code": 0,
//		"message": "删除成功",
//	})
//}

type FeedbackController struct {

}

func (this *FeedbackController) FeedbackList(c *gin.Context) {

	list := new(service.FeedbackService).FindList()

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"message": "success",
		"result": list,
	})
}

func (this *FeedbackController) DelFeedback(c *gin.Context) {
	field := c.PostForm("field")

	if err := new(service.FeedbackService).Delete(field); err != nil {
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

func (this *FeedbackController) DelAllFeedback(c *gin.Context) {
	if err := new(service.FeedbackService).DelAll(); err != nil {
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