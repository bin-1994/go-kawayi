package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-biu/app/admin/user/model"
	"go-biu/library/database/orm"
	"go-biu/utils"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func FetchAllVip(c *gin.Context) {
	offset, length := utils.GetPage(c.Query("page"), c.Query("pagesize"))

	var (
		vipList []model.Vip
		total   int64
	)

	db := orm.Db
	keyword := c.Query("keyword")
	if keyword != "" {
		db = db.Where("order_number LIKE ?", fmt.Sprintf("%%%s%%", keyword))
	}

	db.Limit(length).Offset(offset).Order("vid DESC").Find(&vipList)
	db.Table("vip").Count(&total)

	c.JSON(http.StatusOK, map[string]interface{}{
		"code":   0,
		"result": vipList,
		"total":  total,
	})
}

func FetchVip(c *gin.Context) {
	vid := c.Param("vid")
	var vipList model.Vip
	orm.Db.Where("vid = ?", vid).First(&vipList)
	c.JSON(http.StatusOK, map[string]interface{}{
		"code":   0,
		"result": vipList,
	})
}

func EditVip(c *gin.Context) {
	var form model.VipForm

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code":    1,
			"message": "请输入必填字段",
		})
		return
	}

	if err := orm.Db.Model(&form).Updates(&form).Error; err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code":    1,
			"message": "编辑失败",
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "编辑成功",
	})
}

func CheckVip(c *gin.Context) {
	vid := c.Param("vid")

	var vip model.Vip
	orm.Db.Where("vid = ?", vid).First(&vip)
	if 1 == vip.State {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code":    1,
			"message": "已经审核完成，请不要重复操作",
		})
		return
	}

	var user model.User
	now := int(time.Now().Unix())
	orm.Db.Where("user_id = ?", vip.Uid).First(&user)
	dayTime := vip.Day * 86400
	data := make(map[string]interface{})
	if user.VipEndTime > now {
		// 会员没过期，续费
		data["vip_end_time"] = gorm.Expr("vip_end_time + ?", dayTime)
	} else {
		// 会员已过期
		data["vip_start_time"] = now
		data["vip_end_time"] = dayTime + now
	}
	tx := orm.Db.Begin()
	if err := tx.Model(&user).Where("user_id = ?", vip.Uid).Updates(data).Error; err != nil {
		tx.Rollback()
		return
	}

	if err := tx.Model(&vip).Where("vid = ?", vid).Updates(map[string]interface{}{"state": 1, "start_time": now, "end_time": now + dayTime}).Error; err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, map[string]interface{}{
		"code":    0,
		"message": "审核完成",
	})
}

func NotPassVip(c *gin.Context) {
	vid := c.Param("vid")
	var vip model.Vip

	orm.Db.Where("vid = ?", vid).First(&vip)
	if 1 == vip.State {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code":    1,
			"message": "已经审核完成，请不要重复操作",
		})
		return
	}

	orm.Db.Model(&vip).Where("vid = ?", vid).Updates(map[string]interface{}{
		"state": -1,
	})
}
