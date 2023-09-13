package router

import (
	"gin_tarvel_repository/constant"
	"gin_tarvel_repository/exception"
	"github.com/gin-gonic/gin"
)

func MethodNotAllowedRouter(Router *gin.Engine) {
	//405
	Router.NoMethod(
		func(context *gin.Context) {
			error := &exception.ServiceError{
				Message: "Method Not Allow", Code: constant.MethodNotAllowed,
			}
			//错误处理
			context.Error(error)
		},
	)
}
