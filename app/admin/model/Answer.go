package model

type AnswerModel struct {
	Id       int64  `json:"id" form:"id"`
	Question string `json:"question" form:"question"`
	Option1  string `json:"option1" form:"option1"`
	Option2  string `json:"option2" form:"option2"`
	Option3  string `json:"option3" form:"option3"`
	Option4  string `json:"option4" form:"option4"`
	Answer   string `json:"answer" form:"answer"`
}
