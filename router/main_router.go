package router

import (
	"gin_tarvel_repository/handler"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitRouterGroup(Router *gin.Engine) {
	initHandler(Router)
	initSession(Router)

	UserRouter(Router)
	LoginRouter(Router)
	TravelRouter(Router)
	SwaggerRouter(Router)
}

// 初始化中间件
func initHandler(Router *gin.Engine) {
	//全局异常处理中间件
	Router.Use(handler.GlobalErrorHandler())
}

func initSession(Router *gin.Engine) {
	//初始化session
	store := cookie.NewStore([]byte("secret"))
	Router.Use(sessions.Sessions("session", store))
}
