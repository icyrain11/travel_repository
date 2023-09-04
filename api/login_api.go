package api

import (
	"gin_tarvel_repository/requst"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginApi interface {
	LoginByPassword(userLoginRequest requst.UserLoginRequest,
		session sessions.Session) error

	Logout(c *gin.Context)
}
