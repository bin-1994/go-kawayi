package model

import (
	"gorm.io/gorm"
)

type Season struct {
	MediaId uint `json:"media_id" gorm:"primaryKey"`
	SeasonTitle string `json:"title"`
	//VideoCover string `json:"video_cover"`
	BVideoCover string `json:"video_cover"`
	IsFinish int `json:"is_finish"`
	EpIndex string `json:"ep_index"`
	EpTotal int `json:"ep_total"`
	Follows int `json:"follows"`
	SeasonId int `json:"-"`
}

func (Season) TableName () string {
	return "media"
}

func (this *Season) AfterFind(tx *gorm.DB) (err error) {
	this.BVideoCover = IMG_PREFIX + this.BVideoCover
	return
}

