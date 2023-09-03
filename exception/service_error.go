package exception

import (
	"fmt"
	"log"
)

// ServiceError 服务错误结构体
type ServiceError struct {
	Message string
	Code    int
	Request string
}

// 服务异常处理
func (serviceError *ServiceError) Error() string {
	//服务类异常
	logMessage := fmt.Sprintf("Service服务异常错误代码: %d , 错误消息: %s , 调用服务名称: %s",
		serviceError.Code, serviceError.Message, serviceError.Request)
	log.Println(logMessage)

	return logMessage
}
