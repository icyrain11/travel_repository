package api

import (
	"gin_tarvel_repository/request"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginApi interface {
	LoginByPassword(userLoginRequest request.UserLoginRequest,
		session sessions.Session) error

	Logout(c *gin.Context)
}
