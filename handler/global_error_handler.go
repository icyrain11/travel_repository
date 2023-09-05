package handler

import (
	"gin_tarvel_repository/constant"
	"gin_tarvel_repository/exception"
	"gin_tarvel_repository/model/common"

	"github.com/gin-gonic/gin"
)

// GlobalErrorHandler  全局异常处理中间件
func GlobalErrorHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()
		//处理异常
		if error := context.Errors.Last(); error != nil {
			switch error.Err.(type) {
			case *exception.ServiceError:
				{
					error := error.Err.(*exception.ServiceError)
					handlerServiceError(context, error)
					break
				}
			}
		}
		// 防止错误继续传递给下一个中间件或处理函数
		context.Abort()
	}
}

func handlerServiceError(context *gin.Context, serviceError *exception.ServiceError) {
	code := serviceError.Code
	message := serviceError.Message
	switch code {
	case constant.Fail:
		{
			common.FailWithMessage(message, context)
			break
		}
	case constant.UnAuthorized:
		{
			common.FailWithDetail(constant.UnAuthorized, message, map[string]interface{}{}, context)
			break
		}
	default:
		{
			common.FailWithMessage(message, context)
			break
		}
	}

}
