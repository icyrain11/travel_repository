package router

import (
	"gin_tarvel_repository/exception"
	"github.com/gin-gonic/gin"
)

// NotFoundRouter 404参数路由
func NotFoundRouter(Router *gin.Engine) {
	//404
	Router.NoRoute(
		func(context *gin.Context) {
			error := &exception.ServiceError{
				Message: "Not Found", Code: 404,
			}
			//错误处理
			context.Error(error)
		},
	)
}
