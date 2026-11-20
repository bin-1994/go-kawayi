package model

type Danmaku struct {
	Id int64 `json:"id"`
	Uid int64 `json:"uid"`
	MediaId int `json:"media_id"`
	EpId int64 `json:"ep_id"`
	Mode int `json:"mode"`
	Color int `json:"color"`
	Text string `json:"text"`
	Size int `json:"size"`
	Stime int `json:"stime"`
	Ip string `json:"ip"`
	Ctime int64 `json:"ctime"`
}
