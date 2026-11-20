package model

type User struct {
	Uid int64 `json:"uid" gorm:"primaryKey"`
	Role int `json:"role"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Avatar string `json:"avatar"`
	RegIp string `json:"reg_ip"`
	VipStartTime int `json:"vip_start_time"`
	VipEndTime int `json:"vip_end_time"`
	LastLoginIp string `json:"last_login_ip"`
	LastLoginTime int `json:"last_login_time"`
	CreateTime int `json:"create_time"`
	UpdateTime int `json:"update_time"`
	State int `json:"state"`
}

type UserForm struct {
	Uid int64 `form:"uid" gorm:"primaryKey"`
	Email string `form:"email" binding:"required"`
	Phone string `form:"phone"`
	Username string `form:"username"`
	Nickname string `form:"nickname" binding:"required"`
	Avatar string `form:"avatar"`
}

func (UserForm) TableName () string {
	return "user"
}
