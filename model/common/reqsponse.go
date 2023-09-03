package common

import (
	"gin_tarvel_repository/constant"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:code`
	Data    interface{} `json:data`
	Message string      `json:message`
}

func Result(code int, data interface{}, message string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code, data, message,
	})
}

func Success(c *gin.Context) {
	Result(constant.Success, map[string]interface{}{}, "操作成功", c)
}

func SuccessWithMessage(message string, c *gin.Context) {
	Result(constant.Success, map[string]interface{}{}, message, c)
}

func SuccessWithMessageAndData(message string, data interface{}, c *gin.Context) {
	Result(constant.Success, data, message, c)
}

func Fail(c *gin.Context) {
	Result(constant.Fail, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(constant.Fail, map[string]interface{}{}, message, c)
}

func FailWithMessageAndData(message string, data interface{}, c *gin.Context) {
	Result(constant.Fail, data, message, c)
}

func FailWithDetail(code int, message string, data interface{}, c *gin.Context) {
	Result(code, data, message, c)
}
