package model

type Season struct {
	SeasonId int `json:"season_id" form:"season_id" gorm:"primaryKey"`
	Title string `json:"title" form:"title"`
	Keyword string `json:"keyword" form:"keyword"`
	State int `json:"state" form:"state"`
}

func (Season) TableName () string {
	return "media_season"
}
