package main

import (
	"github.com/MohamedAbdelrazeq/gin-video-server/controller"
	"github.com/MohamedAbdelrazeq/gin-video-server/service"
	"github.com/gin-gonic/gin"
)

var (
	VideoService    = service.NewVideoService()
	videoController = controller.NewVideoController(VideoService)
)

func main() {
	router := gin.Default()

	videos := router.Group("/videos")
	{
		videos.GET("/", func(ctx *gin.Context) { videoController.FindAll(ctx) })
		videos.POST("/", func(ctx *gin.Context) { videoController.Save(ctx) })
	}

	router.Run(":8080")
}
