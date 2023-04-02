package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hotbrandon/go-gin-tutorial/entity"
	"github.com/hotbrandon/go-gin-tutorial/service"
)

type IVideoController interface {
	GetAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type VideoController struct {
	service service.VideoService
}

func (c *VideoController) GetAll() []entity.Video {
	return c.service.GetAll()
}

func (c *VideoController) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	err := ctx.BindJSON(&video)
	if err != nil {
		fmt.Println("VideoController Save error:", err)
	}
	c.service.Save(video)
	return video
}

func NewVideoController(svc service.VideoService) VideoController {
	return VideoController{
		service: svc,
	}
}
