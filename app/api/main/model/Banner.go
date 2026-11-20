package model

type Banner struct {
	Id uint `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
	SubTitle string `json:"sub_title"`
	Cover string `json:"cover"`
	DataId uint `json:"data_id"`
}
