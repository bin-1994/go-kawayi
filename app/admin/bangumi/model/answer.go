package model

type Answer struct {
	Qid int `json:"qid" gorm:"primaryKey"`
	Question string `json:"question"`
	Option1 string `json:"option1"`
	Option2 string `json:"option2"`
	Option3 string `json:"option3"`
	Option4 string `json:"option4"`
	Answer string `json:"answer"`
}


type AnswerForm struct {
	Qid int `form:"qid,omitempty" gorm:"primaryKey"`
	Question string `form:"question" binding:"required"`
	Option1 string `form:"option1" binding:"required"`
	Option2 string `form:"option2" binding:"required"`
	Option3 string `form:"option3" binding:"required"`
	Option4 string `form:"option4" binding:"required"`
	Answer string `form:"answer" binding:"required"`
}

func (AnswerForm) TableName() string {
	return "answer"
}
