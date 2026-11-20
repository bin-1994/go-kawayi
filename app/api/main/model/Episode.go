package model

import (
	"fmt"
	"go-biu/utils"
	"gorm.io/gorm"
)

type EpisodeModel struct {
	//Vid         string `json:"vid"`
	EpId        int    `json:"ep_id" gorm:"primaryKey"`
	MediaId     int    `json:"media_id"`
	EpIndex     string `json:"ep_index"`
	TitleFormat string `json:"title_format"`
	Title       string `json:"title"`
	Vurl        string `json:"vurl"`
	BackupVurl  string `json:"backup_vurl"`
	CreateTime  uint   `json:"ctime"`
	Index       int    `json:"index"`
}

func (this *EpisodeModel) AfterFind(tx *gorm.DB) (err error) {
	this.TitleFormat = this.EpIndex
	if utils.IsNumber(this.EpIndex) {
		this.TitleFormat = fmt.Sprintf("第%s话", this.EpIndex)
	}
	//this.Vid = utils.EncodeHashId([]uint64{uint64(this.MediaId), uint64(this.EpId)})
	return
}
