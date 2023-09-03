package router

import (
	"gin_tarvel_repository/handler"
	"github.com/gin-gonic/gin"
)

func InitRouterGroup(Router *gin.Engine) {
	initHandler(Router)
	UserRouter(Router)
	LoginRouter(Router)
	TravelRouter(Router)
	//初始化swagger
	SwaggerRouter(Router)
}

// 初始化中间件
func initHandler(Router *gin.Engine) {
	//全局异常处理中间件
	Router.Use(handler.GlobalExceptionHandler())
}
