package controller

import (
	"gin_tarvel_repository/exception"
	"gin_tarvel_repository/model/common"
	"gin_tarvel_repository/requst"
	"gin_tarvel_repository/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LoginByPassWord(context *gin.Context) {
	//获取登录参数
	userLoginRequest := requst.UserLoginRequest{}
	fullPath := context.FullPath()

	if error := context.ShouldBind(&userLoginRequest); error != nil {
		error := &exception.ServiceError{
			Message: "表单参数错误", Code: 500, Request: fullPath,
		}
		context.Error(error)
		return
	}

	//调用service
	loginService := service.LoginService{}
	session := sessions.Default(context)
	if error := loginService.LoginByPassword(userLoginRequest, session); error != nil {
		context.Error(error)
		return
	}

	//返回结果
	common.Success(context)
}

func Logout(c *gin.Context) {

}
