package model

type Author struct {
	Uid      uint   `json:"uid,omitempty" gorm:"primaryKey"`
	Nickname string `json:"nickname,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
}

type UserModel struct {
	Uid        int64  `json:"uid"`
	Nickname   string `json:"nickname"`
	Email      string `json:"email"`
	Password   string `json:"-"`
	Avatar     string `json:"avatar"`
	RegIp      string `json:"-"`
	CreateTime int64  `json:"-"`
	State      int    `json:"-"`
}

type RegCode struct {
	Ip   string `json:"ip"`
	Code string `json:"code"`
}

type UserProgress struct {
	MediaId     int `json:"media_id"`
	EpId        int `json:"ep_id"`
	FollowState int `json:"follow_state"`
	Progress    int `json:"progress"`
}

type UserReview struct {
	ReviewId   int64  `json:"review_id"`
	Score      int    `json:"score"`
	Content    string `json:"content"`
	IsCoin     bool   `json:"is_coin"`
	UpdateTime int64  `json:"update_time"`
}

type UserState struct {
	Progress UserProgress `json:"progress"`
	Review   UserReview   `json:"review"`
}
