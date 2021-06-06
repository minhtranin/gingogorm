package ctl

import (
	"work/gingo/entity"
	"work/gingo/svc"

	"github.com/gin-gonic/gin"
)

type VideoCtl interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type ctl struct {
	service svc.VideoService
}

func New(svc svc.VideoService) VideoCtl {
	return &ctl{
		service: svc,
	}
}

func (c *ctl) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (c *ctl) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}
