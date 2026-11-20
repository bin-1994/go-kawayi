package model

type FeedbackModel struct {
	MediaId    int    `form:"media_id" json:"media_id"`
	EpId       int    `form:"ep_id" json:"ep_id"`
	Url        string `form:"url" json:"url"`
	Reason     string `form:"reason" json:"reason"`
	Message    string `form:"message" json:"message"`
	Contact    string `form:"contact" json:"contact"`
	Ip         string `form:"ip" json:"ip"`
	UserAgent  string `form:"user_agent" json:"user_agent"`
	CreateTime int64  `form:"create_time" json:"create_time"`
}

func (FeedbackModel) TableName() string {
	return "feedback"
}
