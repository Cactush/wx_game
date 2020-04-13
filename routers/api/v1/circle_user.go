package v1

import (
	"github.com/Cactush/go-gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCircleUser(c *gin.Context) {
	var userId int = 1
	circleUser := models.GetCircleUser(userId)
	c.JSON(http.StatusOK, gin.H{
		"data": circleUser,
	})
}
