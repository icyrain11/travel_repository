package api

import "gin_tarvel_repository/request"

// UserApi 用户api接口
type UserApi interface {
	AddUser(userAddRequest request.UserAddRequest)

	GetUserById(int int)
}
