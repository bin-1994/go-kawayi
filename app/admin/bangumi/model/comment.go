package model

type Comment struct {
	Cid int64 `json:"cid"`
	Uid int64 `json:"uid"`
	ReplyUid int64 `json:"reply_uid"`
	Root int64 `json:"root"`
	Parent int64 `json:"parent"`
	EpId int64 `json:"ep_id"`
	Plat int `json:"plat"`
	Message string `json:"message"`
	Rcount int `json:"rcount"`
	Like int `json:"like"`
	Hate int `json:"hate"`
	Ip string `json:"ip"`
	Ctime int64 `json:"ctime"`
}
