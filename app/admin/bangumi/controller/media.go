//package controller
//
//import (
//	"github.com/gin-gonic/gin"
//	"go-biu/app/admin/bangumi/model"
//	"go-biu/app/admin/bangumi/service"
//	"go-biu/library/database/orm"
//	"go-biu/utils"
//	"net/http"
//	"strconv"
//)
//
//func FetchAllMedia(c *gin.Context) {
//
//	offset, length := utils.GetPage(c.Query("page"), c.Query("pagesize"))
//
//	var (
//		media []model.Media
//		total int64
//	)
//	db := orm.Db.Table("media")
//
//	keyword := c.Query("keyword")
//	if keyword != "" {
//		db = db.Where("MATCH ( `title`, `origin_name`, `keyword` ) AGAINST (? IN BOOLEAN MODE)", "+"+keyword+"+")
//	}
//	is_new := c.Query("is_new")
//	if is_new != "0" {
//		db = db.Where("is_new = ?", is_new)
//	}
//	is_limit := c.Query("is_limit")
//	if is_limit != "0" {
//		db = db.Where("is_limit = ?", is_limit)
//	}
//	version := c.Query("version")
//	if version != "0" {
//		db = db.Where("version = ?", version)
//	}
//
//	db.Count(&total).Limit(length).Offset(offset).Order("media_id DESC").Find(&media)
//
//	c.JSON(http.StatusOK, map[string]interface{}{
//		"code":   0,
//		"result": media,
//		"total":  total,
//	})
//}
//
//func FetchMedia(c *gin.Context) {
//	mediaId := c.Param("media_id")
//
//	var media model.Media
//	orm.Db.Where("media_id = ?", mediaId).First(&media)
//
//	var styles []int
//	orm.Db.Table("media_style").Where("media_id = ?", mediaId).Pluck("style_id", &styles)
//	media.Styles = styles
//
//	c.JSON(http.StatusOK, map[string]interface{}{
//		"code":   0,
//		"result": media,
//	})
//}
//
//func AddMedia(c *gin.Context) {
//	var form model.MediaForm
//	c.ShouldBind(&form)
//
//	if form.Title == "" {
//		c.JSON(http.StatusOK, map[string]interface{}{
//			"code":    1,
//			"message": "请输入标题",
//		})
//		return
//	}
//
//	if err := orm.Db.Create(&form).Error; err != nil {
//		c.JSON(http.StatusOK, map[string]interface{}{
//			"code":    1,
//			"message": "添加失败",
//		})
//		return
//	}
//
//	// 添加风格
//	styles := c.PostFormArray("styles[]")
//	for _, val := range styles {
//		if val != "" {
//			style := model.MediaStyle{MediaId: form.MediaId, StyleId: val}
//			orm.Db.Create(&style)
//		}
//	}
//
//	c.JSON(http.StatusOK, map[string]interface{}{
//		"code":    0,
//		"message": "添加成功",
//	})
//}
//
//func EditMedia(c *gin.Context) {
//	var form model.MediaForm
//	c.ShouldBind(&form)
//	if form.Title == "" {
//		c.JSON(http.StatusOK, map[string]interface{}{
//			"code":    1,
//			"message": "请输入标题",
//		})
//		return
//	}
//
//	if err := orm.Db.Model(&form).Updates(&form).Error; err != nil {
//		c.JSON(http.StatusOK, map[string]interface{}{
//			"code":    1,
//			"message": "编辑失败",
//		})
//		return
//	}
//
//	// 删除再添加风格
//	orm.Db.Where("media_id = ?", form.MediaId).Delete(&model.MediaStyle{})
//	styles := c.PostFormArray("styles[]")
//	for _, val := range styles {
//		if val != "" {
//			style := model.MediaStyle{MediaId: form.MediaId, StyleId: val}
//			orm.Db.Create(&style)
//		}
//	}
//
//	c.JSON(http.StatusOK, map[string]interface{}{
//		"code":    0,
//		"message": "编辑成功",
//	})
//}
//
//func FetchTimeline(c *gin.Context) {
//	var media []model.Media
//
//	orm.Db.Table("media").Where("is_new = 1").Order("pub_time ASC").Find(&media)
//
//	c.JSON(http.StatusOK, map[string]interface{}{
//		"code":    0,
//		"message": "success",
//		"result":  media,
//	})
//}
//
//type MediaController struct {
//}
//
//func (this *MediaController) DeleteCache(c *gin.Context) {
//	id := c.PostForm("media_id")
//
//	new(service.MediaService).DeleteCaceh(id)
//
//	c.JSON(http.StatusOK, map[string]interface{}{
//		"code":    0,
//		"message": "success",
//	})
//}
//
//func (this *MediaController) UpdateTime(c *gin.Context) {
//	media_id := c.PostForm("media_id")
//	id, _ := strconv.Atoi(media_id)
//	new(service.MediaService).UpdateTime(id)
//
//	c.JSON(http.StatusOK, map[string]interface{}{
//		"code":    0,
//		"message": "success",
//	})
//}
