package service

import (
	"gin_tarvel_repository/request"
)

// UserService 实现接口
type UserService struct {
}

func (userService UserService) AddUser(userAddRequest request.UserAddRequest) {
	//email := userAddRequest.Email
}

func (userService UserService) GetUserById(id int) {

}
