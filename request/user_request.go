package request

type UserQueryRequest struct {
}

type UserLoginRequest struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required" `
}

type UserAddRequest struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required" `
	NickName string `json:"nickName"  binding:"required" `
}
