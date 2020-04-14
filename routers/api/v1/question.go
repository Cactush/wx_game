package v1

import (
	"github.com/Cactush/go-gin/models"
	"github.com/Cactush/go-gin/pkg/logging"
	"github.com/gin-gonic/gin"
	"net/http"
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
	}{question,option_dict}
	c.JSON(http.StatusOK, gin.H{
		"data": temp_dict,
	})
}
