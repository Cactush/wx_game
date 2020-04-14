package jwt

import (
	"github.com/Cactush/go-gin/models"
	"github.com/Cactush/go-gin/pkg/e"
	"github.com/Cactush/go-gin/pkg/logging"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		token := c.GetHeader("Authorization")
		code = e.SUCCESS

		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			var user_token = models.Usertoken{}
			logging.Info(token)
			models.Db.Where("`key`=?", token).First(&user_token)
			if user_token != (models.Usertoken{}) {
				logging.Info(token)
				var circleuser = models.Circleuser{}
				models.Db.Where("user_id=?", user_token.UserId).First(&circleuser)
				if circleuser == (models.Circleuser{}) {
					code = e.ERROR
				} else {
					c.Keys = make(map[string]interface{})
					c.Keys["user"] = &circleuser
				}
			} else {
				code = e.ERROR
				logging.Info(user_token)
			}
		}
		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}

}
