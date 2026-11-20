package model

type AnswerModel struct {
	Qid      int    `json:"qid" gorm:"primaryKey"`
	Question string `json:"question"`
	Option1  string `json:"option1"`
	Option2  string `json:"option2"`
	Option3  string `json:"option3"`
	Option4  string `json:"option4"`
	Answer   string `json:"answer"`
}

type AnswerCacheModel struct {
	Timestamp int           `json:"timestamp"`
	List      []AnswerModel `json:"list"`
}
