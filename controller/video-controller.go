package controller

import (
	"net/http"

	"github.com/MohamedAbdelrazeq/gin-video-server/entity"
	"github.com/MohamedAbdelrazeq/gin-video-server/helper"
	"github.com/MohamedAbdelrazeq/gin-video-server/service"
	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll(ctx *gin.Context)
	Save(ctx *gin.Context)
}

type videoController struct {
	service service.VideoService
}

func NewVideoController(service service.VideoService) VideoController {
	return &videoController{
		service: service,
	}
}

func (c *videoController) FindAll(ctx *gin.Context) {
	videos, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"videos": videos,
	})
}

func (c *videoController) Save(ctx *gin.Context) {
	var video entity.Video

	// BINDING & VALIDATION
	if err := ctx.BindJSON(&video); err != nil {
		helper.InvokeError(ctx, "error while binding json")
		return
	}

	// BUSINESS LOGIC
	if _, err := c.service.Save(video); err != nil {
		helper.InvokeError(ctx, "error while saving video")
		return
	}

	// RESPONSE
	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "Video is saved successfully",
		"video":   video,
	})
}
