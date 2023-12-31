package main

import (
	"io"
	"os"
	"poc/controller"
	"poc/middleware"
	"poc/service"

	"github.com/gin-gonic/gin"
)

var(
	videoService service.VideoService = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput(){
	f,_ := os.Create("gin.log")
	gin.DefaultWriter := io.MultiWriter(f,os.Stdout)
}


func main() {

	setupLogOutput()
	server := gin.New()

	server.Use(gin.Recovery(), middleware.Logger())

	server.GET("/videos",func (ctx *gin.Context)  {
		ctx.JSON(200,videoController.FindAll())
	})

	server.POST("/videos",func (ctx *gin.Context)  {
		ctx.JSON(200,videoController.Save(ctx))
	})

	server.Run(":8080")
}