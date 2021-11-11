package main

import (
	"myapp/controller"
	"myapp/service"

	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {

	// Creates a gin router with default middleware
	router := gin.Default()

	// A handler for GET request on /example
	router.GET("/videos", func(ctx *gin.Context) {

		ctx.JSON(200, videoController.FindAll())

	}) // gin.H is a shortcut for map[string]interface{}
	router.POST("/videos", func(ctx *gin.Context) {

		ctx.JSON(200, videoController.Save(ctx))
	})
	router.Run(":8080") // listen and serve on port 8080
}
