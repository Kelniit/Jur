package router

import (
	"Jur/controller"

	"github.com/gin-gonic/gin"
)

func SampleRouter(route *gin.Engine) {
	// Sample Router
	sample := route.Group("/sample")
	{
		sample.GET("/GetAll", controller.GetAll)
		sample.GET("/GetSample/:SampleID", controller.GetSample)
		sample.POST("/Create", controller.CreateSample)
	}
}
