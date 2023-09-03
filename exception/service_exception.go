package exception

// ServiceException 服务类型错误结构体
type ServiceException struct {
	Message string
	Code    int
	Request string
}

func (e *ServiceException) ServiceError() string {
	return e.Message
}
