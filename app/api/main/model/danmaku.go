package model

import (
	"gorm.io/gorm"
	"html"
)

type DanmakuModel struct {
	Id int64 `json:"dm_id" gorm:"primaryKey"`
	Uid int64 `json:"uid" form:"uid"`
	MediaId int `json:"-" form:"media_id" binding:"required"`
	EpId int64 `json:"-" form:"ep_id" binding:"required"`
	Mode int `json:"mode" form:"mode" binding:"required"`
	Color int `json:"color" form:"color" gorm:"default:16777215"`
	Text string `json:"text" form:"text" binding:"required"`
	Size int `json:"size" form:"size" gorm:"default:25"`
	Stime int `json:"stime" form:"stime" binding:"required"`
	Ip string `json:"-"`
	Ctime int64 `json:"ctime" gorm:"autoCreateTime"`
}

func (DanmakuModel) TableName () string {
	return "danmaku"
}

//func (this *DanmakuModel) BeforeCreate(tx *gorm.DB) (err error) {
//	this.Stime = this.Stime * 1000
//	return
//}

func (this *DanmakuModel) AfterFind (tx *gorm.DB) (err error) {
	// 过滤
	//item.Message = "555"
	this.Text = html.EscapeString(this.Text)
	//this.Stime = this.Stime / 1000
	return
}
