package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint `gorm:"primaryKey;default:auto_random()"`
	Email    string
	Password string
	NickName string
}

func (User) TableName() string {
	return "t_user"
}
