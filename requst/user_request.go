package requst

type UserQueryRequest struct {
}

// UserLoginRequest 用户登录结构体
type UserLoginRequest struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required" `
}
