package router

import (
	"gin_tarvel_repository/controller"
	"github.com/gin-gonic/gin"
)

func UserRouter(Router *gin.Engine) {
	UserRouter := Router.Group("user")
	{

		UserRouter.POST("/", controller.AddUser)

		UserRouter.GET("/:id", controller.GetUserById)

	}
}
