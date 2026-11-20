package model

import "gorm.io/gorm"

type FollowModel struct {
	Uid         int64  `json:"uid" gorm:"<-:create"`
	MediaId     int    `json:"media_id" form:"media_id"`
	EpId        int64  `json:"ep_id" form:"ep_id"`
	EpIndex     string `json:"ep_index" form:"ep_index"`
	FollowState int    `json:"follow_state" form:"follow_state"`
	Progress    int    `json:"progress" form:"progress"`
	CreateTime  int64  `json:"create_time"`
	LastTime    int64  `json:"last_time"`
}

type BangumiFollowModel struct {
	MediaId     int    `json:"media_id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Cover       string `json:"cover"`
	VideoCover  string `json:"video_cover"`
	IsFinish    int    `json:"is_finish"`
	Story       string `json:"story"`
	EpId        int64  `json:"ep_id"`
	WatchEp     string `json:"watch_ep"`
	EpIndex     string `json:"ep_index"`
	EpTotal     string `json:"ep_total"`
	Progress    int    `json:"progress"`
	UpdateTime  int64  `json:"update_time"`
	LastTime    int64  `json:"last_time"`
	NotWatch    bool   `json:"not_watch"`
	FollowState int    `json:"follow_state"`
}

func (this *BangumiFollowModel) AfterFind(tx *gorm.DB) error {
	if this.UpdateTime > this.LastTime {
		this.NotWatch = true
	}
	this.Cover = IMG_PREFIX + this.Cover
	this.VideoCover = IMG_PREFIX + this.VideoCover
	return nil
}
