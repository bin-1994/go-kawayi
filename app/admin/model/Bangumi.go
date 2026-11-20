package model

type BangumiModel struct {
	Id          uint64 `json:"id" form:"id"`
	Category    int    `json:"category" form:"category"`
	Version     int    `json:"version" form:"version"`
	Area        int    `json:"area" form:"area"`
	IsFinish    int    `json:"is_finish" form:"is_finish"`
	IsNew       int    `json:"is_new" form:"is_new"`
	IsLimit     int    `json:"is_limit" form:"is_limit"`
	IsJump      int    `json:"is_jump" form:"is_jump"`
	IsHide      int    `json:"is_hide" form:"is_hide"`
	Level       int    `json:"level" form:"level"`
	Title       int    `json:"title" form:"title"`
	OriginName  string `json:"origin_name" form:"origin_name"`
	Keyword     string `json:"keyword" form:"keyword"`
	Subtitle    string `json:"subtitle" form:"subtitle"`
	Tags        string `json:"tags" form:"tags"`
	Week        int    `json:"week" form:"week"`
	OnAir       string `json:"on_air" form:"on_air"`
	PubTime     int64  `json:"pub_time" form:"pub_time"`
	PubString   string `json:"pub_string" form:"pub_string"`
	Tips        string `json:"tips" form:"tips"`
	Url         string `json:"url" form:"url"`
	Story       string `json:"story" form:"story"`
	Cover       string `json:"cover" form:"cover"`
	VideoCover  string `json:"video_cover" form:"video_cover"`
	Staff       string `json:"staff" form:"staff"`
	Cast        string `json:"cast" form:"cast"`
	EpId        uint64 `json:"ep_id" form:"ep_id"`
	EpTitle     string `json:"ep_title" form:"ep_title"`
	EpTotal     string `json:"ep_total" form:"ep_total"`
	SeriesId    uint64 `json:"series_id" form:"series_id"`
	SeriesTitle string `json:"series_title" form:"series_title"`
	CreateTime  int64  `json:"create_time" form:"create_time"`
	UpdateTime  int64  `json:"update_time" form:"update_time"`
	State       int    `json:"state" form:"state"`
}

type BangumiResponse struct {
	BangumiModel
	Styles []int `json:"styles"`
}

type BangumiQuery struct {
	Category int    `form:"category" validate:"number"`
	Version  int    `form:"version" validate:"number"`
	Area     int    `form:"area" validate:"number"`
	State    int    `form:"state" validate:"number"`
	Month    int    `form:"month" validate:"number"`
	Year     string `form:"year"`
	Style    int    `form:"style" validate:"number"`
	Order    int    `form:"order" validate:"number"`
	Sort     int    `form:"sort" validate:"number"`
	IsNew    int    `form:"is_new" validate:"number"`
	Keyword  string `form:"keyword"`
}
