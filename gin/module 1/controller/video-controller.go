package controller

import (
	"myapp/entity"

	"github.com/gin-gonic/gin"
)

type VideoController interface entity.Video {
	Save(ctx *gin.Context) 
	FindAll() []entity.Video
}
type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return controller{
		service: service,
	}
}
func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()

}
func (c *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
}
