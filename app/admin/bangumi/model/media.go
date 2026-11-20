package model

type Media struct {
	MediaId uint `json:"media_id" gorm:"primaryKey"`
	Type uint `json:"type"`
	Version uint `json:"version"`
	Area uint `json:"area"`
	IsFinish uint `json:"is_finish"`
	IsNew uint `json:"is_new"`
	IsLimit uint `json:"is_limit"`
	IsJump uint `json:"is_jump"`
	//IsHide uint `json:"is_hide"`
	Bd uint `json:"bd"`
	//Hd uint `json:"hd"`
	Title string `json:"title"`
	OriginName string `json:"origin_name"`
	Keyword string `json:"keyword"`
	Week int `json:"week"`
	OnAir string `json:"on_air"`
	PubTime int `json:"pub_time"`
	PubString string `json:"pub_string"`
	Duration int `json:"duration"`
	Tips string `json:"tips"`
	Download string `json:"download"`
	Url string `json:"url"`
	Story string `json:"story"`
	Cover string `json:"cover"`
	VideoCover string `json:"video_cover"`
	BCover string `json:"b_cover"`
	BVideoCover string `json:"b_video_cover"`
	//SmallCover string `json:"small_cover"`
	Tags string `json:"tags"`
	Level int `json:"level"`
	Subtitle string `json:"subtitle"`
	Staff string `json:"staff"`
	Cast string `json:"cast"`
	EpId uint `json:"ep_id"`
	EpIndex string `json:"ep_index"`
	EpTotal uint `json:"ep_total"`
	SeasonId int `json:"season_id"`
	SeasonTitle string `json:"season_title"`
	Score float32 `json:"score"`
	ScoreTotal uint `json:"score_total"`
	Follows uint `json:"follows"`
	Views uint `json:"views"`
	Comments uint `json:"comments"`
	State uint `json:"state"`
	UpdateTime int64 `json:"update_time"`

	Styles []int `json:"styles" gorm:"-"`
}

type MediaForm struct {
	MediaId uint `form:"media_id,omitempty" gorm:"primaryKey"`
	Title string `form:"title"`
	OriginName string `form:"origin_name"`
	SeasonId int `form:"season_id"`
	SeasonTitle string `form:"season_title"`
	EpIndex *string `form:"ep_index"`
	EpTotal *int `form:"ep_total"`
	EpId *int `form:"ep_id"`
	//Duration *int `form:"duration"`
	PubTime int `form:"pub_time"`
	OnAir *string `form:"on_air"`
	PubString *string `form:"pub_string"`
	Week *int `form:"week"`
	Tips *string `form:"tips"`
	Type *int `form:"type"`
	Version *int `form:"version"`
	Area *int `form:"area"`
	IsFinish *int `form:"is_finish"`
	Bd *int `form:"bd"`
	//Hd *int `form:"hd"`
	IsJump *int `form:"is_jump"`
	IsLimit *int `form:"is_limit"`
	Tags *string `form:"tags"`
	Level *int `form:"level"`
	Cover *string `form:"cover"`
	VideoCover *string `form:"video_cover"`
	BCover *string `form:"b_cover"`
	BVideoCover *string `form:"b_video_cover"`
	//SmallCover string `form:"small_cover"`
	IsNew *int `form:"is_new"`
	Url *string `form:"url"`
	Download *string `form:"download"`
	Keyword *string `form:"keyword"`
	Subtitle *string `form:"subtitle"`
	Staff *string `form:"staff"`
	Cast *string `form:"cast"`
	Story *string `form:"story"`
	State *int `form:"state"`
}

func (MediaForm) TableName () string {
	return "media"
}

type MediaEpRela struct {
	MediaId uint `json:"media_id" gorm:"primaryKey"`
	Title string `json:"title"`
	EpList []Episode `json:"ep_list" gorm:"foreignKey:MediaId"`
}

func (MediaEpRela) TableName () string {
	return "media"
}
