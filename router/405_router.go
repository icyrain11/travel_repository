package router

import (
	"gin_tarvel_repository/constant"
	"gin_tarvel_repository/exception"
	"github.com/gin-gonic/gin"
)

func MethodNotAllowedRouter(Router *gin.Engine) {
	//405
	Router.NoRoute(
		func(context *gin.Context) {
			// 这里只是演示，不要在生产环境中直接返回HTML代码
			error := &exception.ServiceError{
				Message: "Router Not Found", Code: constant.MethodNotAllowed,
			}
			//错误处理
			context.Error(error)
		},
	)
}
