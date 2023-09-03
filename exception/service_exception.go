package exception

import (
	"fmt"
)

// ServiceException 服务类型错误结构体
type ServiceException struct {
	Message string
	Code    int
	Request string
}

// 服务异常处理
func (serviceException *ServiceException) Error() string {
	//服务类异常
	return fmt.Sprintf("Service服务异常错误代码: %d , 错误消息: %s , 调用服务名称: %s",
		serviceException.Code, serviceException.Message, serviceException.Request)
}
