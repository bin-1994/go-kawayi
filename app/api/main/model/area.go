package model

type Area struct {
	Id uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	CountryId uint `json:"-"`
	state uint `json:"-"`
} 
