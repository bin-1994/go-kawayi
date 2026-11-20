package model

type User struct {
	Uid int64 `json:"uid" gorm:"primaryKey"`
	Role int `json:"-"`
	Username string `json:"-"`
	Email string `json:"email"`
	RegIp string `json:"reg_ip"`
	Password string `json:"password,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Avatar string `json:"avatar,omitempty"`
	CreateTime int64 `json:"create_time"`
	State int `json:"state,omitempty" gorm:"default:1"`
}

type RegCode struct {
	Ip string `json:"ip"`
	Code string `json:"code"`
}
