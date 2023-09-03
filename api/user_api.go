package api

import "github.com/gin-gonic/gin"

// UserApi 用户api接口
type UserApi interface {
	AddUser(c *gin.Context)

	GetUserById(c *gin.Context)
}
