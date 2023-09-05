package exception

import (
	"fmt"
	"log"
)

// ValidError 参数校验错误
type ValidError struct {
	Message string
	Code    int
	Request string
}

// TODO 利用反射实现参数骄校验
func (validError ValidError) Error() string {
	logMessage := fmt.Sprintf("Service服务异常错误代码: %d , 错误消息: %s , 调用服务名称: %s",
		validError.Code, validError.Message, validError.Request)
	log.Println(logMessage)
	return logMessage
}
