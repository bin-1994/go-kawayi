package model

type BannerModel struct {
	Id int `json:"id" form:"id,omitempty" gorm:"primaryKey"`
	Title string `json:"title" form:"title"`
	Cover string `json:"cover" form:"cover"`
	DataId int `json:"data_id" form:"data_id"`
	State int `json:"state" form:"state"`
}

func (BannerModel) TableName () string {
	return "banner"
}