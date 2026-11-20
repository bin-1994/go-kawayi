package model

type FeedbackModel struct {
	Id int `json:"id" gorm:"primaryKeys`
	MediaId int `json:"media_id"`
	EpId int `json:"ep_id"`
	Reason string `json:"reason"`
	Message string `json:"message"`
	Contact string `json:"contact"`
	CreateTime int `json:"create_time"`
	Ip string `json:"ip"`
	Ua string `json:"ua"`
	State int `json:"state"`
}

func (FeedbackModel) TableName() string {
	return "feedback"
}


type FeedbackVideo struct {
	MediaId int `json:"media_id"`
	EpId int64 `json:"ep_id"`
	Title string `json:"title"`
	Url string `json:"url"`
	Reason string `json:"reason"`
	Message string `json:"message"`
	Contact string `json:"contact"`
	Ip string `json:"ip"`
	Ua string `json:"ua"`
	Ctime int64 `json:"ctime"`
}