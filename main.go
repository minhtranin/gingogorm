package main

import (
	"work/gingo/ctl"
	"work/gingo/svc"

	"github.com/gin-gonic/gin"
)

var (
	videoService svc.VideoService = svc.New()
	videoCtl     ctl.VideoCtl     = ctl.New(videoService)
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	server := gin.Default()
	server.GET("findAll", func(ctx *gin.Context) {
		ctx.JSON(200, videoCtl.FindAll())
	})
	server.GET("save", func(ctx *gin.Context) {
		ctx.JSON(200, videoCtl.Save(ctx))
	})
	server.Run(":8888")
}
