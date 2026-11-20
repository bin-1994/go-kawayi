package model

type StyleModel struct {
	Id     int64  `form:"id" json:"id"`
	Title  string `form:"title" json:"title"`
	Letter string `form:"letter" json:"letter"`
	State  int    `form:"state" json:"state"`
}

type BangumiStyleModel struct {
	BangumiId int64 `json:"bangumi_id"`
	StyleId   int64 `json:"style_id"`
}
