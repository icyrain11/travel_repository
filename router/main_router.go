package router

import (
	"gin_tarvel_repository/handler"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitRouterGroup(Router *gin.Engine) {

	//中间件
	{
		initHandler(Router)
		initSession(Router)
	}

	//路由
	{
		UserRouter(Router)
		LoginRouter(Router)
		TravelRouter(Router)
	}

	//特殊路由
	{
		NotFoundRouter(Router)
		MethodNotAllowedRouter(Router)
	}

	//swagger
	SwaggerRouter(Router)
}

// 初始化中间件
func initHandler(Router *gin.Engine) {
	//全局异常处理中间件
	Router.Use(handler.GlobalErrorHandler())
	Router.Use(handler.CorsHandler())
}

func initSession(Router *gin.Engine) {
	//初始化session
	store := cookie.NewStore([]byte("secret"))
	Router.Use(sessions.Sessions("session", store))
}
