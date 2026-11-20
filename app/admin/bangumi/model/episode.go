package model

import (
	"gorm.io/gorm"
	"time"
)

type Episode struct {
	EpId int `json:"ep_id" gorm:"primaryKey"`
	MediaId int `json:"media_id"`
	Title string `json:"title"`
	EpIndex string `json:"ep_index"`
	JumpUrl string `json:"jump_url"`
	Version int `json:"version"`
	//Vid string `json:"vid"`
	//DirKey string `json:"dir_key"`
	//PdirKey string `json:"pdir_key"`
	Vurl string `json:"vurl"`
	BackupVurl string `json:"backup_vurl"`
	Source string `json:"source"`
	Index int `json:"index"`
	State bool `json:"state"`
}

type EpForm struct {
	EpId int `form:"ep_id,omitempty" gorm:"primaryKey"`
	MediaId int `form:"media_id"`
	Title string `form:"title"`
	EpIndex string `form:"ep_index"`
	JumpUrl string `form:"jump_url"`
	Source string `form:"source"`
	Version int `form:"version"`
	//Vid string `form:"vid"`
	//DirKey string `form:"dir_key"`
	//PdirKey string `form:"pdir_key"`
	Vurl string `form:"vurl"`
	BackupVurl string `form:"backup_vurl"`
	Index int `form:"index"`
	CreateTime int64 `form:"create_time"`
	UpdateTime int64 `form:"update_time"`
	State *bool `form:"state" gorm:"default:1"`
}

func (EpForm) TableName () string {
	return "episode"
}

func (this *EpForm) BeforeCreate (tx *gorm.DB) (err error) {
	time := time.Now()
	this.CreateTime = time.Unix()
	this.UpdateTime = time.Unix()
	return
}
