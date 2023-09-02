package router

import "github.com/gin-gonic/gin"

func UserRouter(Router *gin.Engine) {
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("list", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "userList",
			})
		})
	}
}
