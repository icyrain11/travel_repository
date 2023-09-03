package router

import (
	"gin_tarvel_repository/api"
	"github.com/gin-gonic/gin"
)

func LoginRouter(Router *gin.Engine) {
	TravelRouter := Router.Group("login")
	{
		TravelRouter.POST("/", api.LoginByPassword)

		TravelRouter.POST("/logout", api.Logout)
	}
}
