package router

import (
	"gin_tarvel_repository/api"
	"github.com/gin-gonic/gin"
)

func UserRouter(Router *gin.Engine) {
	UserRouter := Router.Group("user")
	{

		UserRouter.POST("/", api.AddUser)

		UserRouter.GET("/:id", api.GetUserById)

	}
}
