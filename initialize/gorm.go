package initialize

import (
	"fmt"
	"gin_tarvel_repository/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func Gorm() *gorm.DB {

	gormConfig := config.InitGormConfig()
	url := gormConfig.Url
	port := gormConfig.Port
	passWord := gormConfig.PassWord
	userName := gormConfig.UserName
	dataBase := gormConfig.DataBase
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", userName, passWord, url, port, dataBase)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Open DabaBase Error: %v", err)
	} else {
		log.Println("init db success")
	}

	return db
}
