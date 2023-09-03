package router

import (
	"gin_tarvel_repository/controller"
	"github.com/gin-gonic/gin"
)

func LoginRouter(Router *gin.Engine) {
	LoginRouter := Router.Group("login")
	{
		LoginRouter.POST("/", controller.LoginByPassWord)

		LoginRouter.POST("/logout", controller.Logout)
	}
}
