package model

import "gorm.io/gorm"

type Media struct {
	MediaId int `json:"media_id" gorm:"primaryKey"`
	SeasonId int `json:"series_id"`
	Title string `json:"title"`
	Cover string `json:"cover"`
	VideoCover string `json:"video_cover"`
	IsFinish int `json:"is_finish"`
	Story string `json:"story"`
	EpId int `json:"ep_id"`
	WatchEp string `json:"watch_ep"`
	EpIndex string `json:"ep_index"`
	EpTotal int `json:"ep_total"`
	Progress int `json:"progress"`
	UpdateTime int `json:"-"`
	LastTime int `json:"-"`
	NotWatch bool `json:"not_watch"`
	FollowState int `json:"follow_state"`
}

const IMG_PREFIX = "https://static.kawayi.cc"

func (this *Media) AfterFind(tx *gorm.DB) (err error) {
	if this.UpdateTime > this.LastTime {
		this.NotWatch = true
	}
	this.Cover = IMG_PREFIX + this.Cover
	this.VideoCover = IMG_PREFIX + this.VideoCover
	return
}

type MediaRela struct {
	MediaId int `json:"media_id" gorm:"primaryKey"`
	Title string `json:"title"`
	Cover string `json:"cover"`
}

func (MediaRela) TableName () string {
	return "media"
}
