package model

type HlsModel struct {
	Id int64 `json:"id" gorm:"primary_key"`
	MediaId int `json:"media_id"`
	Name string `json:"name"`
	Dir string `json:"dir"`
	Path string `json:"path"`
	Size int64 `json:"size"`
	Md5 string `json:"md5"`
	Sha1 string `json:"sha1"`
	CreateTime int64 `json:"create_time"`
}

func (HlsModel) TableName() string {
	return "hls"
}
