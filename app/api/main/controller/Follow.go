package controller

import (
	"github.com/gin-gonic/gin"
	"go-biu/app/api/main/model"
	"go-biu/app/api/main/service"
	"go-biu/utils"
	"strconv"
)

type FollowController struct {
	Service service.IFollowService
}

func NewFollowController(service service.IFollowService) *FollowController {
	return &FollowController{Service: service}
}

func (controller *FollowController) Follow(ctx *gin.Context) {
	uid, _ := ctx.Get("uid")
	resp := utils.NewResponse(ctx)
	var follow model.FollowModel
	if err := ctx.ShouldBind(&follow); err != nil {
		resp.Error("数据错误")
		return
	}
	follow.Uid = uid.(int64)
	if err := controller.Service.Create(follow); err != nil {
		resp.Error("追番失败")
		return
	}

	resp.Success(nil)
}

func (controller *FollowController) UnFollow(ctx *gin.Context) {
	uid, _ := ctx.Get("uid")
	resp := utils.NewResponse(ctx)

	var follow model.FollowModel
	if err := ctx.ShouldBind(&follow); err != nil {
		resp.Error("数据错误")
		return
	}
	follow.Uid = uid.(int64)
	if err := controller.Service.Delete(follow); err != nil {
		resp.Error("追番失败")
		return
	}

	resp.Success(nil)
}

func (controller *FollowController) Update(ctx *gin.Context) {
	uid, _ := ctx.Get("uid")
	resp := utils.NewResponse(ctx)
	var follow model.FollowModel
	if err := ctx.ShouldBind(&follow); err != nil {
		resp.Error("数据错误")
		return
	}
	follow.Uid = uid.(int64)
	if err := controller.Service.Update(follow); err != nil {
		resp.Error("更新数据失败")
		return
	}

	resp.Success(nil)
}

func (controller *FollowController) FindAll(ctx *gin.Context) {
	uid, _ := ctx.Get("uid")
	state, _ := strconv.Atoi(ctx.Query("state"))

	page := utils.Paging{Page: ctx.DefaultQuery("page", "1"), Size: ctx.DefaultQuery("size", "15")}
	offset, limit := page.SetPage()

	resp := utils.NewResponse(ctx)

	list, total := controller.Service.FindAll(uid.(int64), state, offset, limit)

	resp.List(list, offset, limit, total)
}

//func Follow(c *gin.Context) {
//	uid, _ := c.Get("uid")
//	var p model.FollowParams
//	if err := c.ShouldBind(&p); err != nil {
//		c.JSON(http.StatusOK, gin.H{
//			"code": 1,
//			"message": "数据错误",
//		})
//		return
//	}
//
//	if err := new(service.FollowService).Create(int64(uid.(float64)), &p); err != nil {
//		c.JSON(http.StatusOK, gin.H{
//			"code": 1,
//			"message": err.Error(),
//		})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"code": 0,
//		"message": "追番成功",
//	})
//}

//func UnFollow(c *gin.Context) {
//	uid, _ := c.Get("uid")
//	var p model.FollowParams
//	if err := c.ShouldBind(&p); err != nil {
//		c.JSON(http.StatusOK, gin.H{
//			"code": 1,
//			"message": "数据错误",
//		})
//		return
//	}
//
//	if err := new(service.FollowService).Delete(int64(uid.(float64)), p.SeriesId); err != nil {
//		c.JSON(http.StatusOK, gin.H{
//			"code": 1,
//			"message": err.Error(),
//		})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"code": 0,
//		"message": "哼！竟然真的抛弃了人家(> <)",
//	})
//}

//func RestFollowNum(c *gin.Context) {
//	series_id := c.PostForm("series_id")
//	id, _ := strconv.Atoi(series_id)
//	if err := new(service.FollowService).Rest(id); err != nil {
//		c.JSON(http.StatusOK, gin.H{
//			"code": 1,
//			"message": err.Error(),
//		})
//		return
//	}
//	c.JSON(http.StatusOK, gin.H{
//		"code": 0,
//		"message": "重置成功",
//	})
//}

//func Progress (c *gin.Context) {
//uid := c.GetInt("uid")
//uid, exists := c.Get("uid")
//if !exists {
//	return
//}
//media_id := c.PostForm("media_id")
//ep_id := c.PostForm("ep_id")
//ep_index := c.PostForm("ep_index")
//ep_order := c.PostForm("ep_order")
//
//if result := orm.Db.Table("media_follow").Where(map[string]interface{}{"media_id": media_id, "uid": uid}).First(&model.MediaFollow{}); result.RowsAffected > 0 {
//	c.JSON(http.StatusOK, map[string]interface{}{
//		"code": 1,
//		"message": "数据不存在",
//	})
//	return
//}
//
//if err := orm.Db.Exec("UPDATE `media_follow` SET `ep_id` = ?, `ep_order` = ?, `episode` = ?, `last_time` = ? WHERE `uid` = ? AND `media_id` = ?", ep_id, ep_order, ep_index, time.Now().Unix(), uid, media_id).Error; err != nil {
//	c.JSON(http.StatusOK, map[string]interface{}{
//		"code": 1,
//		"message": "保存进度失败",
//	})
//	return
//}
//
//c.JSON(http.StatusOK, map[string]interface{}{
//	"code": 0,
//	"message": "保存进度成功",
//})
//	uid, _ := c.Get("uid")
//	var p model.FollowParams
//	if err := c.ShouldBind(&p); err != nil {
//		c.JSON(http.StatusOK, gin.H{
//			"code": 1,
//			"message": "数据错误",
//		})
//		return
//	}
//
//	if err := new(service.FollowService).UpdateProgress(int64(uid.(float64)), &p); err != nil {
//		c.JSON(http.StatusOK, map[string]interface{}{
//			"code": 1,
//			"message": err.Error(),
//		})
//		return
//	}
//
//	c.JSON(http.StatusOK, map[string]interface{}{
//		"code": 0,
//		"message": "保存进度成功",
//	})
//}

//func UpdateFollowState (c *gin.Context) {
//	//uid := c.GetInt("uid")
//	uid, exists := c.Get("uid")
//	if !exists {
//		return
//	}
//	media_id := c.PostForm("media_id")
//	state := c.PostForm("state")
//
//	//if state != "1" || state != "2" || state != "3" {
//	//	c.JSON(http.StatusOK, map[string]interface{}{
//	//		"code": 1,
//	//		"message": "数据错误",
//	//	})
//	//	return
//	//}
//
//	if err := orm.Db.Exec("UPDATE `media_follow` SET `follow_state`= ? WHERE uid = ? AND media_id = ?", state, uid, media_id).Error; err != nil {
//		c.JSON(http.StatusOK, map[string]interface{}{
//			"code": 1,
//			"message": "更新状态失败",
//		})
//		return
//	}
//
//	c.JSON(http.StatusOK, map[string]interface{}{
//		"code": 0,
//		"message": "更新状态成功",
//	})
//}
