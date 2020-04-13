package models

type Circleuser struct {
	UserId   int    `json:"user_id"`
	OpenId   string `json:"open_id"`
	Avatar   string `json:"avatar"`
	NickName string `json:"nick_name"`
}

func GetCircleUser(user_id int) (circleuser Circleuser) {
	Db.Where("user_id=?", user_id).First(&circleuser)
	return
}
