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
			// 这里只是演示，不要在生产环境中直接返回HTML代码
			error := &exception.ServiceError{
				Message: "Method Not Allow", Code: 404,
			}
			//错误处理
			context.Error(error)
		},
	)
}
