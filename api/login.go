package api

import (
	"fmt"
	"gin_tarvel_repository/model/common"
	"gin_tarvel_repository/requst"
	"github.com/gin-gonic/gin"
)

func LoginByPassword(c *gin.Context) {
	//获取登录参数
	userLoginRequest := requst.UserLoginRequest{}

	if err := c.ShouldBind(&userLoginRequest); err != nil {
		//TODO 统一异常处理
		fmt.Print("格式异常")
	}

	//用户id
	//id := 1

	//返回结果
	common.Success(c)
}

func Logout(c *gin.Context) {

}
