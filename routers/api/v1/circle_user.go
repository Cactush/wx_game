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
	c.JSON(http.StatusOK, gin.H{
		"data": circle_user,
	})
}
