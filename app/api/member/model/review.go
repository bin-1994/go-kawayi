package model

type Review struct {
	ReviewId int `json:"review_id" gorm:"primaryKey"`
	MediaId int `json:"-"`
	Score int `json:"score"`
	Content string `json:"content"`
	UpdateTime int `json:"update_time"`
	Media MediaRela `json:"media" gorm:"foreignKey:Id;associationForeignKey:MediaId"`
}
