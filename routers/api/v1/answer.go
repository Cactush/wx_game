package v1

import (
	"encoding/json"
	"github.com/Cactush/go-gin/models"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

func Answer(c *gin.Context) {
	circleUser := c.Keys["user"].(*models.Circleuser)
	var params map[string]interface{}
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fuck"})
	}
	questionerId := com.StrTo(params["questioner_id"].(string)).MustInt()
	answer := params["answer"].([]interface{})
	var correctAnswer []models.Circleuser2question
	models.Db.Where("user_id=?", questionerId).Order("position").Find(&correctAnswer)
	var correctAnswerList []int
	for _, value := range correctAnswer {
		correctAnswerList = append(correctAnswerList, value.Option)
	}
	count := 0
	for index, value := range answer {
		if int(value.(float64)) == correctAnswerList[index] {
			count += 1
		}
	}
	suitability := int(count * 100 / 3)
	userAnswer, _ := json.Marshal(params["answer"])
	var answerRecord = models.Answerrecord{
		QuestionerId: questionerId,
		AnswererId:   circleUser.UserId,
		Suitability:  suitability,
		Answer:       string(userAnswer),
	}
	models.Db.Create(&answerRecord)
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

