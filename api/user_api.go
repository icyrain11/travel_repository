package api

import "gin_tarvel_repository/requst"

// UserApi 用户api接口
type UserApi interface {
	AddUser(userAddRequest requst.UserAddRequest)

	GetUserById(int int)
}
