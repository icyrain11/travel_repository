package handler

import "github.com/gin-gonic/gin"

// TODO GlobalExceptionHandler 全局异常处理
func GlobalExceptionHandler(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			//c.JSON()
		}
	}()

	c.Next()
}
