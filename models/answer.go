package models

type Answerrecord struct {
	Id           int    `json:"id"`
	QuestionerId int    `json:"questioner_id"`
	AnswererId   int    `json:"answerer_id"`
	Suitability  int    `json:"suitability"`
	Answer       string `json:"answer"`
}
