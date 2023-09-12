package controller

import (
	"gin_tarvel_repository/exception"
	"gin_tarvel_repository/model/common"
	"gin_tarvel_repository/request"
	"gin_tarvel_repository/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// LoginByPassWord godoc
// @Summary      Login by pass word
// @Description  login
// @Tags         login
// @Accept       json
// @Produce      json
// @Param		 userLoginRequest
// @Router       /login/ [post]
func LoginByPassWord(context *gin.Context) {
	//获取登录参数
	userLoginRequest := request.UserLoginRequest{}

	if error := context.ShouldBind(&userLoginRequest); error != nil {
		//构造异常
		error := &exception.ServiceError{
			Message: "表单参数错误", Code: 500, Request: "/login/",
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

func Logout(context *gin.Context) {

}
