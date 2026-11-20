package model

type Style struct {
	StyleId int `json:"style_id" form:"style_id"`
	Title string `json:"title" form:"title" binding:"required"`
	Letter string `json:"letter" form:"letter" binding:"required"`
	State int `json:"state" form:"state"`
}
