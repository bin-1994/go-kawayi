package model

type StyleModel struct {
	StyleId int    `json:"style_id" gorm:"primaryKey"`
	Title   string `json:"title"`
}
