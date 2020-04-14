package v1

import (
	"github.com/Cactush/go-gin/models"
	"github.com/Cactush/go-gin/pkg/logging"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCircleUser(c *gin.Context) {

	circle_user := c.Keys["user"].(*models.Circleuser)
	logging.Info(circle_user.KeyWord())
	temp_dict := struct {
		models.Circleuser
		KeyWords      []string `json:"key_words"`
		IsCard        bool     `json:"is_card"`
		IsSetQuestion bool     `json:"is_set_question"`
	}{*circle_user,
		circle_user.KeyWord(),
		circle_user.IsCard(),
		circle_user.IsSetQuestion()}
	c.JSON(http.StatusOK, gin.H{
		"data": temp_dict,
	})
}
