package service

import (
	"gin_tarvel_repository/requst"
)

// UserService 实现接口
type UserService struct {
}

func (userService UserService) AddUser(userAddRequest requst.UserAddRequest) {
	//email := userAddRequest.Email
}

func (userService UserService) GetUserById(id int) {

}
