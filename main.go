package main

import (
	"net/http"
	"work/gingo/ctl"
	"work/gingo/svc"

	"github.com/gin-gonic/gin"
)

var (
	videoService svc.VideoService = svc.New()
	videoCtl     ctl.VideoCtl     = ctl.New(videoService)
	mnS          svc.Manipulate   = svc.NewManuPulate()
	mnCtl        ctl.Manipulate   = ctl.NewManuCtl(mnS)
)

func main() {
	gin.SetMode(gin.DebugMode)
	server := gin.Default()
	server.GET("findAll", func(ctx *gin.Context) {
		ctx.JSON(200, videoCtl.FindAll())
	})
	server.POST("save", func(ctx *gin.Context) {
		err := videoCtl.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		}
		// ctx.JSON(200, )
	})
	server.POST("saveMN", func(ctx *gin.Context) {
		err := mnCtl.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		}
		// ctx.JSON(200, )
	})
	server.Run(":8888")
}
