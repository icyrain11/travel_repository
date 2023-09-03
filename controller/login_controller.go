package controller

import (
	"gin_tarvel_repository/exception"
	"gin_tarvel_repository/model/common"
	"gin_tarvel_repository/requst"
	"gin_tarvel_repository/service"
	"github.com/gin-gonic/gin"
)

func LoginByPassWord(context *gin.Context) {
	//获取登录参数
	userLoginRequest := requst.UserLoginRequest{}

	if error := context.ShouldBind(&userLoginRequest); error != nil {
		fullPath := context.FullPath()
		error := &exception.ServiceException{
			Message: "表单参数错误", Code: 500, Request: fullPath,
		}
		context.Error(error)
		return
	}

	//调用service
	loginService := service.LoginService{}
	//登录
	loginService.LoginByPassword(userLoginRequest)

	//返回结果
	common.Success(context)
}

func Logout(c *gin.Context) {

}
