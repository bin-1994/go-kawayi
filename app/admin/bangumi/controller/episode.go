package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-biu/app/admin/bangumi/model"
	"go-biu/app/admin/bangumi/service"
	"go-biu/library/database/cache"
	"go-biu/library/database/orm"
	"go-biu/utils"
	"net/http"
	"strconv"
	"time"
)

type EpisodeController struct {

}

//func FetchMaxEp (c *gin.Context) {
//	media_id := c.Param("media_id")
//	version := c.Query("version")
//
//	var ep model.Episode
//	orm.Db.Where("media_id = ? AND version = ?", media_id, version).Order("`index` DESC").First(&ep)
//
//	c.JSON(http.StatusOK, map[string]interface{}{
//		"code": 0,
//		"result": ep,
//	})
//}
//
//func FetchMediaToEp (c *gin.Context) {
//	media_id := c.Param("media_id")
//
//	offset, length := utils.GetPage(c.Query("page"), c.Query("pagesize"))
//
//	var (
//		media model.MediaEpRela
//		count int64
//	)
//	orm.Db.Where("media_id = ?", media_id).Preload("EpList", func(db *gorm.DB) *gorm.DB {
//		return db.Order("`index` DESC, ep_id DESC").Offset(offset).Limit(length)
//	}).First(&media)
//
//	orm.Db.Table("episode").Where("media_id = ?", media_id).Count(&count)
//
//	c.JSON(http.StatusOK, map[string]interface{}{
//		"code": 0,
//		"message": "success",
//		"result": media,
//		"total": count,
//	})
//}

func FetchAllEp(c *gin.Context) {
	offset, length := utils.GetPage(c.Query("page"), c.Query("pagesize"))

	media_id := c.Query("media_id")

	var (
		ep []model.Episode
		count int64
	)

	db := orm.Db

	if (media_id != "") {
		db = db.Where("media_id = ?", media_id)
	}

	db.Order("ep_id DESC").Offset(offset).Limit(length).Find(&ep)
	db.Table("episode").Count(&count)

	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"message": "success",
		"result": ep,
		"total": count,
	})
}

//func AddEp (c *gin.Context) {
//	var form model.EpForm
//	c.ShouldBind(&form)
//
//	if err := orm.Db.Create(&form).Error; err != nil {
//		c.JSON(http.StatusOK, map[string]interface{}{
//			"code": 1,
//			"message": "添加失败",
//		})
//		return
//	}
//
//	if is_update := c.PostForm("is_update"); is_update == "1" {
//		orm.Db.Exec("UPDATE `media` SET `ep_index` = ?, `ep_id` = ?, `update_time` = ?, `tips` = '' WHERE `media_id` = ?", form.EpIndex, form.EpId, time.Now().Unix(), form.MediaId)
//	}
//
//	cache.Delete(fmt.Sprintf("media_%s", form.MediaId))
//	cache.Delete("timeline")
//	cache.Delete("timeline_notlogin")
//
//	c.JSON(http.StatusOK, map[string]interface{}{
//		"code": 0,
//		"message": "添加成功",
//	})
//}

//func EditEp (c *gin.Context) {
//	var form model.EpForm
//	c.ShouldBind(&form)
//
//	if err := orm.Db.Model(&form).Updates(&form).Error; err != nil {
//		c.JSON(http.StatusOK, map[string]interface{}{
//			"code": 1,
//			"message": "编辑失败",
//		})
//		return
//	}
//
//	c.JSON(http.StatusOK, map[string]interface{}{
//		"code": 0,
//		"message": "编辑成功",
//	})
//}

func AddEps (c *gin.Context) {
	media_id := c.PostForm("media_id")
	source := c.PostForm("source")
	title := c.PostFormArray("title[]")
	ep_index := c.PostFormArray("ep_index[]")
	//version := c.PostFormArray("version[]")
	jump_url := c.PostFormArray("jump_url[]")
	index := c.PostFormArray("index[]")
	vurl := c.PostFormArray("vurl[]")

	var epModel model.EpForm
	epAfter := make(map[string]interface{})
	db := orm.Db.Begin()
	for key, _ := range ep_index {
		if result := orm.Db.Where("media_id = ? AND `index` = ?", media_id, index[key]).First(&epModel); result.RowsAffected > 0 {
			continue
		}
		media_id, _ := strconv.Atoi(media_id)
		i, _ := strconv.Atoi(index[key])
		//v, _ := strconv.Atoi(version[key])
		epItem := &model.EpForm{
			MediaId: media_id,
			Source: source,
			Title: title[key],
			EpIndex: ep_index[key],
			Version: 1,
			JumpUrl: jump_url[key],
			Vurl: vurl[key],
			Index: i,
		}
		if err := db.Table("episode").Create(epItem).Error; err != nil {
			db.Rollback()
			c.JSON(http.StatusOK, map[string]interface{}{
				"code": 1,
				"message": fmt.Sprintf("操作失败, 第%d个", key + 1),
			})
			return
		}
		epAfter["ep_id"] = epItem.EpId
		epAfter["ep_index"] = epItem.EpIndex
	}
	epAfter["update_time"] = time.Now().Unix()
	db.Table("media").Where("media_id = ?", media_id).Updates(epAfter)
	db.Commit()
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"message": "操作成功",
	})
}

func ReplaceEp (c *gin.Context) {
	media_id := c.PostForm("media_id")
	source := c.PostForm("source")
	isBackup := c.PostForm("isBackup")
	//title := c.PostFormArray("title[]")
	ep_index := c.PostFormArray("ep_index[]")
	jump_url := c.PostFormArray("jump_url[]")
	index := c.PostFormArray("index[]")
	vurl := c.PostFormArray("vurl[]")
	var epModel model.EpForm
	epModel.Source = source
	for key, _ := range index {
		if isBackup != "" {
			epModel.BackupVurl = vurl[key]
		} else if vurl[key] != "" {
			epModel.Vurl = vurl[key]
		} else if jump_url[key] != "" {
			epModel.JumpUrl = jump_url[key]
		}
		orm.Db.Table("episode").Where("media_id = ? AND ep_index = ?", media_id, ep_index[key]).Updates(&epModel)
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"message": "操作成功",
	})
}

func (this *EpisodeController) Get(c *gin.Context) {
	id := c.Param("media_id")
	var p utils.Paging
	c.ShouldBind(&p)
	offset, limit := p.SetPage(p.Page, p.Pagesize)
	info, total := new(service.EpisodeService).Get(id, offset, limit)

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"message": "success",
		"result": info,
		"total": total,
	})
}

func (this *EpisodeController) Create(c *gin.Context) {
	var form model.EpForm
	c.ShouldBind(&form)

	is_update := c.PostForm("is_update")

	if err := new(service.EpisodeService).Create(form, is_update); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"message": err.Error(),
		})
		return
	}

	cache.Delete(fmt.Sprintf("media_%s", form.MediaId))
	cache.Delete("timeline")

	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"message": "添加成功",
	})
}

func (this *EpisodeController) Update(c *gin.Context) {
	var form model.EpForm
	c.ShouldBind(&form)

	if err := new(service.EpisodeService).Update(form); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"message": "编辑成功",
	})
}

func (this *EpisodeController) Delete(c *gin.Context) {
	id := c.PostForm("ep_id")

	if err := new(service.EpisodeService).Delete(id); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"message": "删除成功",
	})
}

func (this *EpisodeController) CreateEp(c *gin.Context) {
	media_id := c.PostForm("media_id")
	source := c.PostForm("source")
	title := c.PostFormArray("title[]")
	ep_index := c.PostFormArray("ep_index[]")
	//version := c.PostFormArray("version[]")
	jump_url := c.PostFormArray("jump_url[]")
	index := c.PostFormArray("index[]")
	vurl := c.PostFormArray("vurl[]")

	if err := new(service.EpisodeService).CreateEp(media_id, source, title, ep_index, jump_url, index, vurl); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"message": "添加成功",
	})
}