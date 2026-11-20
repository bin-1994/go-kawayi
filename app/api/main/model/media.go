package model

//type BaseMedia struct {
//	MediaId      int    `json:"media_id" gorm:"primaryKey"`
//	EpId         int    `json:"ep_id"`
//	Type         int    `json:"type"`
//	Version      int    `json:"version"`
//	Area         uint   `json:"-"`
//	IsFinish     *uint  `json:"is_finish"`
//	Title        string `json:"title"`
//	SeriesFollow int    `json:"series_follow"`
//	Follows      uint   `json:"follows"`
//	Cover        string `json:"cover"`
//	VideoCover   string `json:"video_cover"`
//	UpdateTime   int64  `json:"update_time"`
//	EpIndex      string `json:"ep_index"`
//	EpTotal      uint   `json:"ep_total"`
//	Views        int    `json:"views"`
//}

//func (BaseMedia) TableName() string {
//	return "media"
//}

//const IMG_PREFIX = "https://static.kawayi.cc"

//func (item *BaseMedia) AfterFind(tx *gorm.DB) error {
//
//	item.Cover = IMG_PREFIX + item.Cover
//	item.VideoCover = IMG_PREFIX + item.VideoCover
//	return nil
//}

//type Media struct {
//	BaseMedia
//	//Area uint `json:"area"`
//	SeasonId   int      `json:"season_id"`
//	IsLimit    int      `json:"is_limit"`
//	IsJump     uint     `json:"is_jump"`
//	Keyword    string   `json:"keyword"`
//	OnAir      string   `json:"on_air"`
//	PubTime    int      `json:"pub_time"`
//	PubString  string   `json:"pub_string"`
//	Tips       string   `json:"tips"`
//	Bd         int      `json:"bd"`
//	Download   string   `json:"download"`
//	Url        string   `json:"url"`
//	Story      string   `json:"story"`
//	Staff      string   `json:"staff"`
//	Cast       string   `json:"cast"`
//	Score      *float32 `json:"score"`
//	ScoreTotal *uint    `json:"score_total"`
//	// 关联
//	Episodes []Episode `json:"episodes" gorm:"foreignKey:MediaId"`
//	Areas    Area      `json:"areas" gorm:"foreignKey:CountryId;References:Area"`
//	Seasons  []Season  `json:"seasons" gorm:"foreignKey:SeasonId;References:SeasonId"`
//	Styles   []Style   `json:"styles" gorm:"many2many:media_style;foreignKey:MediaId;joinForeignKey:StyleId;References:StyleId;JoinReferences:MediaId""`
//	Reviews  []Review  `json:"reviews" gorm:"foreignKey:MediaId"`
//	//Relevants []Relevant `json:"relevants"`
//}

//type MediaModel struct {
//	MediaId uint `json:"media_id"`
//	Type    int  `json:"-"`
//}

//type MediaIndex struct {
//	MediaId      int     `json:"media_id"`
//	Type         int     `json:"type"`
//	IsFinish     int     `json:"is_finish"`
//	Title        string  `json:"title"`
//	Subtitle     string  `json:"subtitle"`
//	SeriesFollow int64   `json:"series_follow"`
//	Cover        string  `json:"cover"`
//	EpIndex      string  `json:"ep_index"`
//	EpTotal      int     `json:"ep_total"`
//	Views        int64   `json:"views"`
//	Bd           bool    `json:"bd"`
//	PubTime      int64   `json:"pub_time"`
//	PubString    string  `json:"pub_string"`
//	Score        float64 `json:"score"`
//	SeriesCount  int     `json:"series_count"`
//	UpdateTime   int64   `json:"update_time"`
//}

//func (this *MediaIndex) AfterFind(tx *gorm.DB) (err error) {
//	this.Cover = IMG_PREFIX + this.Cover
//	return
//}

//type TimeLine struct {
//	//BaseMedia
//	MediaId      int    `json:"media_id"`
//	EpId         int64  `json:"ep_id"`
//	Title        string `json:"title"`
//	Cover        string `json:"cover"`
//	VideoCover   string `json:"video_cover"`
//	EpIndex      string `json:"ep_index"`
//	EpTotal      int    `json:"ep_total"`
//	IsFinish     int    `json:"is_finish"`
//	Views        int    `json:"views"`
//	SeriesFollow int64  `json:"series_follow"`
//	Week         int    `json:"week"`
//	UpdateTime   int    `json:"update_time"`
//	New          bool   `json:"new"`
//}

//func (this *TimeLine) AfterFind(tx *gorm.DB) (err error) {
//	diff := time.Now().Unix() - int64(this.UpdateTime)
//	if diff < 86400 {
//		this.New = true
//	}
//	this.Cover = IMG_PREFIX + this.Cover
//	this.VideoCover = IMG_PREFIX + this.VideoCover
//	return
//}

//type Online struct {
//	BaseMedia
//	PubTime   int    `json:"pub_time"`
//	PubString string `json:"pub_string"`
//	Subtitle  string `json:"subtitle"`
//}

//type Relevant struct {
//	MediaId    int    `json:"media_id"`
//	Title      string `json:"title"`
//	Cover      string `json:"cover"`
//	VideoCover string `json:"video_cover"`
//}

//type Media struct {
//	MediaId      int     `json:"media_id" gorm:"primaryKey"`
//	Type         int     `json:"type"`
//	Version      int     `json:"version"`
//	Area         int     `json:"area"`
//	IsFinish     int     `json:"is_finish"`
//	IsNew        bool    `json:"is_new"`
//	IsLimit      int     `json:"is_limit"`
//	IsJump       bool    `json:"is_jump"`
//	Level        int     `json:"level"`
//	Bd           bool    `json:"bd"`
//	Title        string  `json:"title"`
//	OriginName   string  `json:"origin_name"`
//	Keyword      string  `json:"keywords"`
//	Subtitle     string  `json:"subtitle"`
//	Story        string  `json:"story"`
//	Tags         string  `json:"tags"`
//	Week         int     `json:"week"`
//	OnAir        string  `json:"on_air"`
//	PubTime      int64   `json:"pub_time"`
//	PubString    string  `json:"pub_string"`
//	Tips         string  `json:"tips"`
//	Download     string  `json:"website"`
//	Url          string  `json:"url"`
//	Cover        string  `json:"cover"`
//	VideoCover   string  `json:"video_cover"`
//	Staff        string  `json:"staff"`
//	Cast         string  `json:"cast"`
//	EpId         int64   `json:"ep_id"`
//	EpIndex      string  `json:"ep_index"`
//	EpTotal      int     `json:"ep_total"`
//	SeasonId     int     `json:"season_id"`
//	Score        float64 `json:"score"`
//	ScoreTotal   int64   `json:"score_total"`
//	Follows      int64   `json:"follows"`
//	SeriesFollow int64   `json:"series_follow"`
//	Views        int64   `json:"views"`
//}

//func (this *Media) AfterFind(tx *gorm.DB) error {
//	this.Cover = IMG_PREFIX + this.Cover
//	this.VideoCover = IMG_PREFIX + this.VideoCover
//	if this.IsLimit >= 1 {
//		this.Views = 0
//		this.SeriesFollow = 0
//	}
//	return nil
//}

//type MediaDetail struct {
//	MediaId    int           `json:"media_id"`
//	SeasonId   int           `json:"season_id"`
//	Type       int           `json:"type"`
//	Version    int           `json:"version"`
//	Title      string        `json:"title"`
//	Keyword    string        `json:"keywords"`
//	Story      string        `json:"story"`
//	Bd         bool          `json:"bd"`
//	IsJump     bool          `json:"is_jump"`
//	IsLimit    int           `json:"is_limit"`
//	Level      int           `json:"level"`
//	Cover      string        `json:"cover"`
//	VideoCover string        `json:"video_cover"`
//	Url        string        `json:"url"`
//	Download   string        `json:"website"`
//	Staff      string        `json:"staff"`
//	Cast       string        `json:"cast"`
//	Publish    MediaPublish  `json:"publish"`
//	Count      MediaCount    `json:"count"`
//	Rating     MediaRating   `json:"rating"`
//	Areas      Area          `json:"areas"`
//	Seasons    []MediaSeries `json:"seasons"`
//	Related    []TimeLine    `json:"related"`
//	Styles     []Style       `json:"styles"`
//	Reviews    []Review      `json:"reviews"`
//	Episodes   []Episode     `json:"episodes"`
//	//Pv []Episode `json:"pv"`
//	//Sp []Episode `json:"sp"`
//}

//type MediaPublish struct {
//	IsFinish  int    `json:"is_finish"`
//	EpIndex   string `json:"ep_index"`
//	EpTotal   int    `json:"ep_total"`
//	OnAir     string `json:"on_air"`
//	PubTime   int64  `json:"pub_time"`
//	PubString string `json:"pub_string"`
//	Tips      string `json:"tips"`
//}

//type MediaCount struct {
//	Follows      int64 `json:"follows"`
//	SeriesFollow int64 `json:"series_follow"`
//	Views        int64 `json:"views"`
//}

//type MediaSeries struct {
//	MediaId     int    `json:"media_id"`
//	SeasonTitle string `json:"season_title"`
//	VideoCover  string `json:"video_cover"`
//	Follows     int64  `json:"follows"`
//	IsFinish    int    `json:"is_finish"`
//	EpIndex     string `json:"ep_index"`
//	EpTotal     int    `json:"ep_total"`
//}

//func (this *MediaSeries) AfterFind(tx *gorm.DB) (err error) {
//	this.VideoCover = IMG_PREFIX + this.VideoCover
//	return
//}

//type MediaRating struct {
//	Score      float64 `json:"score"`
//	ScoreTotal int64   `json:"total"`
//}

//type UserProgress struct {
//	SeriesId    int   `json:"series_id"`
//	MediaId     int   `json:"media_id"`
//	EpId        int64 `json:"ep_id"`
//	FollowState int   `json:"follow_state"`
//	Progress    int   `json:"progress"`
//}
//
//type UserReview struct {
//	ReviewId   int64  `json:"review_id"`
//	Score      int    `json:"score"`
//	Content    string `json:"content"`
//	IsCoin     bool   `json:"is_coin"`
//	UpdateTime int64  `json:"utime"`
//}
//
//type UserState struct {
//	Progress UserProgress `json:"progress"`
//	Review   UserReview   `json:"review"`
//}

//type MediaParams struct {
//	MediaId int   `form:"media_id"`
//	EpId    int64 `form:"ep_id"`
//}

// 番剧索引
//type MediaIndexParams struct {
//	Page     int    `form:"page"`
//	Pagesize int    `form:"pagesize"`
//	Category int    `form:"category"`
//	Version  int    `form:"version"`
//	State    int    `form:"state"`
//	Area     int    `form:"area"`
//	Month    int    `form:"month"`
//	Year     string `form:"year"`
//	Style    int    `form:"style"`
//	Order    int    `form:"order"`
//	Sort     int    `form:"sort"`
//}

//type Filter struct {
//	Filter []FilterList `json:"filter"`
//	Order  FilterList   `json:"order"`
//}
//
//type FilterList struct {
//	Title string       `json:"title"`
//	Key   string       `json:"key,omitempty"`
//	Items []FilterItem `json:"items"`
//}
//
//type FilterItem struct {
//	Value string `json:"value"`
//	Name  string `json:"name"`
//	Sort  string `json:"sort,omitempty"`
//}
