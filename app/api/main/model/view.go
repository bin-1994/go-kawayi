package model

import "gorm.io/gorm"

type ViewModel struct {
	MediaId      int    `json:"media_id"`
	Title        string `json:"title"`
	Subtitle     string `json:"subtitle"`
	Cover        string `json:"cover"`
	VideoCover   string `json:"video_cover"`
	Views        int64  `json:"views"`
	SeriesFollow int64  `json:"series_follow"`
	EpIndex      string `json:"ep_index"`
	EpTotal      string `json:"ep_total"`
	IsFinish     int    `json:"is_finish"`
}

func (ViewModel) TableName() string {
	return "media"
}

func (this *ViewModel) AfterFind(tx *gorm.DB) error {
	this.Cover = IMG_PREFIX + this.Cover
	this.VideoCover = IMG_PREFIX + this.VideoCover
	return nil
}
