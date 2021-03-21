package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gdarias1987/serverApp/controller"
	"github.com/gdarias1987/serverApp/middlewares"
	"github.com/gdarias1987/serverApp/service"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService       = service.New()
	VideoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogOutput()

	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger(),
		middlewares.BasicAuth(), gindump.Dump())

	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "OK!!",
		})
	})

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		err := VideoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "video added succesfully",
			})
		}
	})

	server.Run(":8080")
	fmt.Println("Server running...")
}
