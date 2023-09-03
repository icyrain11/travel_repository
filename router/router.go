package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouterGroup(Router *gin.Engine) {
	UserRouter(Router)
	LoginRouter(Router)
	TravelRouter(Router)
	//初始化swagger
	SwaggerRouter(Router)
}
