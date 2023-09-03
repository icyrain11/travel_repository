package service

import (
	"fmt"
	"gin_tarvel_repository/requst"
	"github.com/gin-gonic/gin"
)

type LoginService struct {
}

func (loginService LoginService) LoginByPassword(userLoginRequest requst.UserLoginRequest) {
	fmt.Print(userLoginRequest)
}

func (loginService LoginService) Logout(c *gin.Context) {

}
