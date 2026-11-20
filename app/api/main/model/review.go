package model

type Review struct {
	ReviewId int64 `json:"review_id" gorm:"primaryKey"`
	Uid int64 `json:"-"`
	MediaId int `json:"-"`
	Score int `json:"score"`
	Content string `json:"content"`
	IsCoin bool `json:"-"`
	CreateTime int64 `json:"-"`
	UpdateTime int64 `json:"utime"`
	Author Author `json:"author" gorm:"foreignKey:Uid;references:Uid"`
}

type ReviewForm struct {
	ReviewId int64 `json:"review_id" form:"review_id" gorm:"primaryKey;<-:create"`
	Uid int64 `json:"uid" form:"uid" gorm:"<-:create"`
	MediaId int `json:"media_id" form:"media_id" gorm:"<-:create"`
	Score int `json:"score" form:"score"`
	Content string `json:"content" form:"content"`
	CreateTime int64 `json:"create_time" form:"create_time" gorm:"autoCreateTime"`
	UpdateTime int64 `json:"update_time" form:"update_time" gorm:"autoUpdateTime"`
}
