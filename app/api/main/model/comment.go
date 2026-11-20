package model

import (
	"gorm.io/gorm"
	"html"
)

type BaseComment struct {
	Cid      int64  `json:"cid" gorm:"primaryKey"`
	Uid      int64  `json:"-"`
	ReplyUid int64  `json:"-" form:"reply_uid"`
	Root     int64  `json:"root" form:"root"`
	Parent   int64  `json:"parent" form:"parent"`
	EpId     int64  `json:"-" form:"ep_id" binding:"required"`
	Plat     uint   `json:"plat" form:"plat"`
	Message  string `json:"message" form:"message" binding:"required"`
	Rcount   uint   `json:"rcount"`
	Like     uint   `json:"like"`
	Hate     uint   `json:"-"`
	Ip       string `json:"-"`
	Ctime    int64  `json:"ctime"`
}

func (BaseComment) TableName() string {
	return "comment"
}

func (this *BaseComment) AfterFind(tx *gorm.DB) (err error) {
	// 过滤
	//item.Message = "555"
	this.Message = html.EscapeString(this.Message)
	return
}

type ReplyModel struct {
	BaseComment
	At     Author `json:"at,omitempty" gorm:"ForeignKey:Uid;references:ReplyUid"`
	Author Author `json:"member" gorm:"ForeignKey:Uid;references:Uid"`
}

type CommentModel struct {
	BaseComment
	Author Author       `json:"member" gorm:"ForeignKey:Uid;references:Uid"`
	Replys []ReplyModel `json:"replys" gorm:"ForeignKey:Root;references:Cid"`
}
