package main

import (
	"fmt"
	"github.com/Cactush/go-gin/pkg/setting"
	"github.com/Cactush/go-gin/routers"
	"github.com/fvbock/endless"
	"log"
	"syscall"
)

func main() {
	//router := gin.Default()
	//router := routers.InitRouter()
	//router.GET("/test", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "test",
	//	})
	//})
	//s := &http.Server{
	//	Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
	//	Handler:        router,
	//	TLSConfig:      nil,
	//	ReadTimeout:    setting.ReadTimeout,
	//	WriteTimeout:   setting.WriteTimeout,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//s.ListenAndServe()

	// 优雅重启
	endless.DefaultReadTimeOut = setting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1<<20
	endPoint := fmt.Sprintf(":%d",setting.HTTPPort)

	server:=endless.NewServer(endPoint,routers.InitRouter())

	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d",syscall.Getpid())

	}
	err:= server.ListenAndServe()
	if err!=nil{
		log.Printf("server err:%v",err)
	}
}
