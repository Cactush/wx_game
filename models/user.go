package models

import (
	"fmt"
	"github.com/Cactush/go-gin/pkg/logging"
	"github.com/unknwon/com"
	"github.com/wxnacy/wgo/arrays"
	"strings"
)

type Circleuser struct {
	UserId           int    `json:"user_id"`
	OpenId           string `json:"open_id"`
	Avatar           string `json:"avatar"`
	NickName         string `json:"nick_name"`
	WechatAvatar     string `json:"wechat_avatar"`
	Age              string `json:"-"`
	Education        string `json:"-"`
	PlaceOfResidence string `json:"place"`
	Height           int    `json:"-"`
	Gender           string `json:"-"`
	House            string `json:"-"`
	MonthlySalary    string `json:"-"`
	Car              string `json:"-"`
}
type Usertoken struct {
	Key    string `json:"key"`
	UserId int    `json:"user_id"`
}

func GetCircleUser(user_id int) (circleuser Circleuser) {
	Db.Where("user_id=?", user_id).First(&circleuser)
	return
}

func (c *Circleuser) GetTag() []string {
	return make([]string, 1)
}
func (c *Circleuser) GetAge() (age_keyword string) {
	temp_age := c.Age
	defer func() {
		if err := recover(); err != nil {
			age_keyword = ""
			logging.Info("名字关键字出错")
		}
	}()
	age := com.StrTo(temp_age[2:4]).MustInt()
	year := age / 10
	num := age % 10 / 5
	age_keyword = fmt.Sprintf("%d后", (year*10 + num*5))
	return
}

func (c *Circleuser) GetEducation() string {
	education := c.Education
	if arrays.Contains([]string{"本科", "硕士", "博士"}, education) == -1 {
		return ""
	}
	return education
}
func (c*Circleuser)GetAddress()string  {
	return c.PlaceOfResidence

}
func (c *Circleuser) GetHeight() (height_keywork string) {
	if c.Gender == "男" && c.Height > 175 {
		height_keywork = "身高175+"
	}
	if c.Gender == "女" && c.Height > 160 {
		height_keywork = "身高160+"
	}
	return
}

func (c *Circleuser) GetHouse() (house_keywork string) {
	if strings.ContainsAny(c.House, "已购房") {
		house_keywork = "有房"
	}
	return
}

func (c *Circleuser) GetMonthlySalary()(monthly_salary string) {
	if c.MonthlySalary != "" && strings.ContainsAny(c.MonthlySalary, "5千") {
		monthly_salary = "月薪过万"
	}
	return
}

func (c *Circleuser) GetCar()(car string) {
	if c.Car!=""&& strings.ContainsAny(c.Car,"有车"){
		car = c.Car
	}
	return
}

func (c *Circleuser)KeyWord() []string  {
	keyWordList := []string{c.GetAge(),c.GetEducation(),c.GetAddress(),
		c.GetHeight(),c.GetHouse(),c.GetMonthlySalary(),c.GetCar()}
	temp_keyworklist := []string{}
	for _,value := range keyWordList{
		if value!=""{
			temp_keyworklist = append(temp_keyworklist,value)
		}
	}
	return temp_keyworklist[:6]
}
