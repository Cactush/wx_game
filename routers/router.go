package routers

import (
	"github.com/Cactush/go-gin/middleware/jwt"
	"github.com/Cactush/go-gin/pkg/setting"
	"github.com/Cactush/go-gin/routers/api"
	v1 "github.com/Cactush/go-gin/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter()*gin.Engine  {
	r:=gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	r.GET("/auth",api.GetAuth)
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/tags",v1.GetTags)
		apiv1.POST("tags",v1.AddTag)
		apiv1.PUT("/tags/:id",v1.EditTag)
		apiv1.DELETE("/tags/:id",v1.DeleteTag)


		apiv1.GET("/articles",v1.GetArticles)
		apiv1.GET("articles/:id",v1.GetArticle)
		apiv1.POST("/articles",v1.AddArticle)
		apiv1.PUT("/articles/:id",v1.EditArticle)
		apiv1.DELETE("/articles/:id",v1.DeleteArticle)

		apiv1.GET("/circleuser",v1.GetCircleUser)
		apiv1.GET("/question/:id",v1.GetQuestion)
		apiv1.POST("/set_question",v1.SetQuestion)
	}
	return r
}