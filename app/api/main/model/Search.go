package model

import (
	"gorm.io/gorm"
)

type SearchModel struct {
	MediaId   int        `json:"media_id" gorm:"primaryKey"`
	Title     string     `json:"title"`
	Keyword   string     `json:"keyword"`
	Cover     string     `json:"cover"`
	PubTime   int        `json:"pub_time"`
	PubString string     `json:"pub_string"`
	Area      int        `json:"area"`
	Story     string     `json:"story"`
	Type      int        `json:"type"`
	Version   int        `json:"version"`
	Cast      string     `json:"cast"`
	Episodes  []SearchEp `json:"episodes,omitempty" gorm:"foreignKey:MediaId"`
}

func (SearchModel) TableName() string {
	return "media"
}

func (this *SearchModel) AfterFind(tx *gorm.DB) error {
	this.Cover = IMG_PREFIX + this.Cover
	return nil
}

type SearchEp struct {
	//Vid     string `json:"vid"`
	EpId    int    `json:"ep_id" gorm:"primaryKey"`
	MediaId int    `json:"media_id"`
	Title   string `json:"title"`
	EpIndex string `json:"ep_index"`
	Index   string `json:"index"`
}

func (SearchEp) TableName() string {
	return "episode"
}

func (this *SearchEp) AfterFind(tx *gorm.DB) (err error) {
	//this.Vid = utils.EncodeHashId([]uint64{uint64(this.MediaId), uint64(this.EpId)})
	return
}
