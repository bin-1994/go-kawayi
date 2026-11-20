package model

import "gorm.io/gorm"

const IMG_PREFIX = "https://static.kawayi.cc"

type BangumiModel struct {
	MediaId    int    `gorm:"primaryKey"`
	SeasonId   int    `json:"season_id"`
	Title      string `json:"title"`
	Subtitle   string `json:"subtitle"`
	Keyword    string `json:"keyword"`
	Level      int    `json:"level"`
	Cover      string `json:"cover"`
	VideoCover string `json:"video_cover"`
	Staff      string `json:"staff"`
	Cast       string `json:"cast"`
	Story      string `json:"story"`
	IsFinish   int    `json:"is_finish"`
	IsLimit    int    `json:"is_limit"`
	EpIndex    string `json:"ep_index"`
	EpId       int64  `json:"ep_id"`
	EpTotal    int    `json:"ep_total"`
	PubTime    int64  `json:"pub_time"`
	OnAir      string `json:"on_air"`
	PubString  string `json:"pub_string"`
	Tips       string `json:"tips"`
	Follows    uint64 `json:"follows"`
	Views      uint64 `json:"views"`
}

//func (BangumiModel) TableName() string {
//	return "media"
//}

func (this *BangumiModel) AfterFind(tx *gorm.DB) (err error) {
	if this.Cover != "" {
		this.Cover = IMG_PREFIX + this.Cover
	}
	if this.VideoCover != "" {
		this.VideoCover = IMG_PREFIX + this.VideoCover
	}
	return
}

type BangumiListModel struct {
	MediaId    int     `json:"media_id" gorm:"primaryKey"`
	Title      string  `json:"title"`
	Subtitle   string  `json:"subtitle"`
	Week       int     `json:"week"`
	PubTime    int64   `json:"pub_time"`
	PubString  string  `json:"pub_string"`
	Cover      string  `json:"cover"`
	VideoCover string  `json:"video_cover"`
	IsFinish   int     `json:"is_finish"`
	EpId       int64   `json:"ep_id"`
	EpIndex    string  `json:"ep_index"`
	EpTotal    int     `json:"ep_total"`
	Follows    uint64  `json:"follows"`
	Views      uint64  `json:"views"`
	Score      float64 `json:"score,omitempty"`
	UpdateTime int64   `json:"update_time"`
}

func (this *BangumiListModel) AfterFind(tx *gorm.DB) (err error) {
	if this.Cover != "" {
		this.Cover = IMG_PREFIX + this.Cover
	}
	if this.VideoCover != "" {
		this.VideoCover = IMG_PREFIX + this.VideoCover
	}
	return
}

type BangumiInfoModel struct {
	MediaId    int            `json:"media_id"`
	Title      string         `json:"title"`
	Subtitle   string         `json:"sub_title"`
	Keyword    string         `json:"keyword"`
	Level      int            `json:"level"`
	IsLimit    int            `json:"is_limit"`
	Cover      string         `json:"cover"`
	VideoCover string         `json:"video_cover"`
	Staff      string         `json:"staff"`
	Cast       string         `json:"cast"`
	Story      string         `json:"story"`
	Publish    PublishModel   `json:"publish"`
	Count      CountModel     `json:"count"`
	Styles     []StyleModel   `json:"styles"`
	Seasons    []SeasonModel  `json:"seasons"`
	Episode    []EpisodeModel `json:"episodes"`
}

type PublishModel struct {
	IsFinish  int    `json:"is_finish"`
	EpIndex   string `json:"ep_index"`
	EpId      int64  `json:"ep_id"`
	EpTotal   int    `json:"ep_total"`
	PubTime   int64  `json:"pub_time"`
	OnAir     string `json:"on_air"`
	PubString string `json:"pub_string"`
	Tips      string `json:"tips"`
}

type CountModel struct {
	Follows uint64 `json:"follows"`
	Views   uint64 `json:"views"`
}

type SeasonModel struct {
	MediaId     int    `json:"media_id"`
	Title       string `json:"title"`
	SeasonTitle string `json:"season_title"`
}

type BangumiQueryModel struct {
	Category int    `form:"category"`
	Version  int    `form:"version"`
	Area     int    `form:"area"`
	State    int    `form:"state"`
	Month    int    `form:"month"`
	Year     string `form:"year"`
	Style    int    `form:"style"`
	Order    int    `form:"order"`
	Sort     int    `form:"sort"`
	IsNew    int    `form:"is_new"`
	Page     int    `form:"page"`
	Size     int    `form:"size"`
}

type FilterModel struct {
	Filter []FilterList `json:"filter"`
	Order  FilterList   `json:"order"`
}

type FilterList struct {
	Title string       `json:"title"`
	Key   string       `json:"key,omitempty"`
	Items []FilterItem `json:"items"`
}

type FilterItem struct {
	Value string `json:"value"`
	Name  string `json:"name"`
	Sort  string `json:"sort,omitempty"`
}
