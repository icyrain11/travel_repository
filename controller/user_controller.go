package controller

import (
	"gin_tarvel_repository/exception"
	"gin_tarvel_repository/request"
	"gin_tarvel_repository/service"
	"github.com/gin-gonic/gin"
)

// AddUser 接口
func AddUser(context *gin.Context) {
	userAddRequest := request.UserAddRequest{}
	fullPath := context.FullPath()
	if error := context.ShouldBind(&userAddRequest); error != nil {
		error := &exception.ServiceError{
			Message: "表单参数错误", Code: 500, Request: fullPath,
		}
		context.Error(error)
		return
	}
	userService := service.UserService{}
	userService.AddUser(userAddRequest)
}

func GetUserById(c *gin.Context) {

}
