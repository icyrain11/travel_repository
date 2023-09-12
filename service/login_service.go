package service

import (
	"gin_tarvel_repository/constant"
	"gin_tarvel_repository/exception"
	"gin_tarvel_repository/request"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginService struct {
}

func (loginService LoginService) LoginByPassword(userLoginRequest request.UserLoginRequest,
	session sessions.Session) error {

	//返回错误
	var error error

	error = &exception.ServiceError{
		Message: "发生错误了!", Code: constant.Fail, Request: "/login/LoginByPassword",
	}

	id := 1
	session.Set("user", id)
	session.Save()

	return error
}

func (loginService LoginService) Logout(c *gin.Context) {

}
