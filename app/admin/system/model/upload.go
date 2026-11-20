package model

type UploadModel struct {
	Id int `json:"id" gorm:"primaryKey"`
	Dir string `json:"dir"`
	Name string `json:"name"`
	Path string `json:"path"`
	Md5 string `json:"md5"`
	Sha1 string `json:"sha1"`
	Ext string `json:"ext"`
	Size int64 `json:"size"`
	CreateTime int64 `json:"create_time"`
}

func (UploadModel) TableName() string {
	return "upload"
}
