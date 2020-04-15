package routers

import (
	"github.com/Cactush/go-gin/middleware/jwt"
	"github.com/Cactush/go-gin/pkg/setting"
	"github.com/Cactush/go-gin/routers/api"
	v1 "github.com/Cactush/go-gin/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	r.GET("/auth", api.GetAuth)
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{

		apiv1.GET("/circleuser", v1.GetCircleUser)
		apiv1.GET("/question", v1.GetQuestion)
		apiv1.POST("/set_question", v1.SetQuestion)
		apiv1.GET("/questions", v1.GetQuestions)
		apiv1.GET("/question/:id", v1.GetQuestionDetail)

		apiv1.POST("/answer",v1.Answer)
	}
	return r
}
