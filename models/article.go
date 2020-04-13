package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag Tag `json:"tag"`

	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string`json:"modified_by"`
	State int `json:"state"`
}

func (article *Article)BeforeCreate(scope gorm.Scope)error  {
	scope.SetColumn("CreatedOn",time.Now().Unix())
	return nil
}


func (article *Article)BeforeUpdate(scope gorm.Scope)error  {
	scope.SetColumn("ModifiedOn",time.Now().Unix())
	return nil
}

func ExistArticleByID(id int)bool  {
	var article Article
	Db.Select("id").Where("id=?",id).First(&article)

	if article.ID>0{
		return true
	}
	return false
}

func GetArticleTotal(maps interface{})(count int)  {
	Db.Model(&Article{}).Where(maps).Count(&count)
	return
}
func GetArticles(pageNum int, pageSize int,maps interface{}) (article []Article){
	Db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&article)
	return
}

func GetArticle(id int)(article Article)  {
	Db.Where("id=?",id).First(&article)
	Db.Model(&article).Related(&article.Tag)
	return
}

func EditArticle(id int, data interface{})bool  {
	Db.Model(&Article{}).Where("id=?",id).Update(data)
	return true
}

func AddArticle(data map[string]interface{}) bool {
	Db.Create(&Article{
		TagID:      data["tag_id"].(int),
		Title:      data["title"].(string),
		Desc:       data["desc"].(string),
		Content:    data["content"].(string),
		CreatedBy:  data["created_by"].(string),
		State:      data["state"].(int),
	})
	return true
}

func DeleteArticle(id int)bool  {
	Db.Where("id=?",id).Delete(Article{})
	return true
}