package api

import "github.com/gin-gonic/gin"

type LoginApi interface {
	LoginByPassword(c *gin.Context)

	Logout(c *gin.Context)
}
