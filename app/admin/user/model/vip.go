package model

type Vip struct {
	Vid int `json:"vid" gorm:"primaryKey"`
	Uid int `json:"uid"`
	Day int `json:"day"`
	Money int `json:"money"`
	OrderNumber string `json:"order_number"`
	Prove string `json:"prove"`
	Remark string `json:"remark"`
	BuyTime int `json:"buy_time"`
	StartTime int `json:"start_time"`
	EndTime int `json:"end_time"`
	State int `json:"state"`
}

type VipForm struct {
	Vid int `form:"vid,omitempty" gorm:"primaryKey"`
	Uid int `form:"uid" binding:"required"`
	Day int `form:"day" binding:"required"`
	Money int `form:"money" binding:"required"`
	OrderNumber string `form:"order_number"`
	Prove string `form:"prove"`
	Remark string `form:"remark"`
	BuyTime int `form:"buy_time"`
	StartTime int `form:"start_time"`
	EndTime int `form:"end_time"`
	State int `form:"state"`
}

func (VipForm) TableName () string {
	return "vip"
}
