package v1

import (
	"github.com/Cactush/go-gin/models"
	"github.com/Cactush/go-gin/pkg/e"
	"github.com/Cactush/go-gin/pkg/setting"
	"github.com/Cactush/go-gin/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"log"
	"net/http"
)

func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	var data interface{}
	if ! valid.HasErrors() {
		if models.ExistArticleByID(id) {
			data = models.GetArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key:%s,err.message:%s", err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
 

func GetArticles(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
		valid.Range(state, 0, 1, "state").Message("状态只能0或者1")
	}
	var tagId int = -1
	if arg := c.Query("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		maps["tag_id"] = tagId
		valid.Min(tagId, 1, "tag_id").Message("标签id必须大于0")
	}

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		data["lists"] = models.GetArticles(util.GetPage(c), setting.PageSize, maps)
		data["total"] = models.GetArticleTotal(maps)
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key:%s,err.message:%s", err.Key, err.Message)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func AddArticle(c *gin.Context) {
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	createdBy := c.Query("created_by")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()

	valid := validation.Validation{}
	valid.Min("tagId", 1, "tag_id").Message("标签id必须大于0")
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(desc, "desc").Message("desc not empty")
	valid.Required(content, "content").Message("content not empty")
	valid.Required(createdBy, "created_by").Message("created_by not empty")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或者1")

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		if models.ExistTagByID(tagId) {
			data := make(map[string]interface{})
			data["tag_id"] = tagId
			data["title"] = title
			data["desc"] = desc
			data["content"] = content
			data["created_by"] = createdBy
			data["state"] = state

			models.AddArticle(data)
			code = e.SUCCESS
		}else{
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	}else{
		for _,err:= range valid.Errors{
			log.Printf("err.key:%s,err.message:%s",err.Key,err.Message)
		}
	}
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg": e.GetMsg(code),
		"data":make(map[string]interface{}),
	})
}

func EditArticle(c *gin.Context) {
	valid :=validation.Validation{}

	id:=com.StrTo(c.Param("id")).MustInt()
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title :=c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	modifiedBy :=c.Query("modified_by")
	var state int = -1
	if arg:=c.Query("state"); arg!=""{
		state = com.StrTo(arg).MustInt()
		valid.Range(state,0,1,"state").Message("state is 0 or1")
	}
	valid.Min(id,1,"id").Message("id must more than 0")
	valid.MaxSize(title,100,"title").Message("title less than 100")
	valid.MaxSize(desc,255,"desc").Message("desc less than 255")
	valid.MaxSize(content,65535,"content").Message("content less then65535")
	valid.Required(modifiedBy,"modified_by").Message("modifiedBy not empty")
	valid.MaxSize(modifiedBy,100,"modified_by").Message("modifiedBy not empty")

	code:=e.INVALID_PARAMS
	if !valid.HasErrors(){
		if models.ExistTagByID(id){
			if models.ExistTagByID(tagId){
				data := make(map[string]interface{})
				data["tag_id"]=tagId
				data["title"]=title
				data["desc"]=desc
				data["content"]=content
				data["modified_by"]=modifiedBy
				models.EditArticle(id,data)
				code=e.SUCCESS
			}else{
				code = e.ERROR_NOT_EXIST_TAG
			}
		}else{
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	}else{
		for _,err:=range valid.Errors{
			log.Printf("err.key %s,err.message %s",err.Key,err.Message)
		}
	}
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.GetMsg(code),
		"data":make(map[string]string),
	})
}

func DeleteArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	valid :=validation.Validation{}
	valid.Min(id,1,"id").Message("id must more than 0")

	code := e.INVALID_PARAMS
	if !valid.HasErrors(){
		if models.ExistArticleByID(id){
			models.DeleteArticle(id)
			code = e.SUCCESS
		}else{
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	}else{
		for _,err :=range valid.Errors{
			log.Printf("err.key:%s,err.message:%s",err.Key,err.Message)
		}
	}
	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.GetMsg(code),
		"data":make(map[string]string),
	})
}
