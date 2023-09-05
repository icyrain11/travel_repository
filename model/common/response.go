package common

import (
	"gin_tarvel_repository/constant"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:code`
	Data    interface{} `json:data`
	Message string      `json:message`
}

func Result(code int, data interface{}, message string, context *gin.Context) {
	context.JSON(http.StatusOK, Response{
		code, data, message,
	})
}

func Success(context *gin.Context) {
	Result(constant.Success, map[string]interface{}{}, "操作成功", context)
}

func SuccessWithMessage(message string, context *gin.Context) {
	Result(constant.Success, map[string]interface{}{}, message, context)
}

func SuccessWithMessageAndData(message string, data interface{}, context *gin.Context) {
	Result(constant.Success, data, message, context)
}

func Fail(context *gin.Context) {
	Result(constant.Fail, map[string]interface{}{}, "操作失败", context)
}

func FailWithMessage(message string, context *gin.Context) {
	Result(constant.Fail, map[string]interface{}{}, message, context)
}

func FailWithMessageAndData(message string, data interface{}, context *gin.Context) {
	Result(constant.Fail, data, message, context)
}

func FailWithDetail(code int, message string, data interface{}, context *gin.Context) {
	Result(code, data, message, context)
}
