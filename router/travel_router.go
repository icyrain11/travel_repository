package router

import (
	"gin_tarvel_repository/api"
	"github.com/gin-gonic/gin"
)

func TravelRouter(Router *gin.Engine) {
	TravelRouter := Router.Group("travel")
	{
		TravelRouter.POST("/", api.AddTravel)

		TravelRouter.GET("/", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "添加",
			})
		})

		TravelRouter.PUT("/", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "添加",
			})
		})

		TravelRouter.DELETE("/", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "添加",
			})
		})
	}
}
