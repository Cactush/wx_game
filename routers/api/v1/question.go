package v1

import (
	"fmt"
	"github.com/Cactush/go-gin/models"
	"github.com/Cactush/go-gin/pkg/logging"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

type QuestionParams struct {
	Already_selected []int
	Already_replace  []int
}

func GetQuestion(c *gin.Context) {
	circle_user := c.Keys["user"].(*models.Circleuser)
	logging.Info(circle_user.UserId)
	var params QuestionParams
	err := c.BindJSON(&params)
	if err != nil {
		logging.Info(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": 400,})
		return
	}
	var question models.Questionbank
	id_list := params.Already_replace
	models.Db.Where("id not in (?)", id_list).Take(&question)

	option_dict, err := question.GetOption()
	temp_dict := struct {
		models.Questionbank
		Option map[int]string `json:"option"`
	}{question, option_dict}
	c.JSON(http.StatusOK, gin.H{
		"data": temp_dict,
	})
}

type SetQuestionParams struct {
	Question []map[string]int
	Type     int
}

func SetQuestion(c *gin.Context) {
	circleUser := c.Keys["user"].(*models.Circleuser)
	logging.Info(circleUser.UserId)
	var params SetQuestionParams
	err := c.BindJSON(&params)
	if err != nil {
		logging.Info("设置问题参数错误", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": 400})
		return
	}
	if params.Type == 1 {
		err := CreateUserQuestion(models.Db, circleUser.UserId, params)
		logging.Info(err)
	}
	if params.Type == 2 {
		param := params.Question[0]
		record := models.Circleuser2question{}
		models.Db.Where("user_id= ? and position=?", circleUser.UserId, param["position"]).First(&record)
		record.Option = param["option"]
		record.QuestionId = param["id"]
		models.Db.Save(&record)
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func CreateUserQuestion(db *gorm.DB, user_id int, params SetQuestionParams) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_id = ?", user_id).Delete(&(models.Circleuser2question{})).Error; err != nil {

			return err
		}
		for index, value := range params.Question {
			record := models.Circleuser2question{
				UserId:      user_id,
				QuestionId:  value["id"],
				Option:      value["option"],
				Position:    index + 3,
				CreatedTime: fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05")),
			}
			if err := tx.Create(&record).Error; err != nil {

				return err
			}
		}
		tx.Commit()
		return nil
	})

}
