package models

type Circleuser2question struct {
	Id          int    `json:"id"`
	UserId      int    `json:"user_id"`
	QuestionId  int    `json:"question_id"`
	Option      int    `json:"option"`
	Position    int    `json:"position"`
	CreatedTime string `json:"created_time"`
}
