package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Mysql struct {
	Url      string `config:"url"`
	Port     string `config:"port"`
	UserName string `config:"username"`
	PassWord string `config:"password"`
	DataBase string `config:"database"`
}

func InitGormConfig() Mysql {
	file, err := os.ReadFile("config.yaml")

	if err != nil {
		log.Fatalf("Failed to read YAML file: %v", err)
	}

	fmt.Println("yaml open content: \n", string(file))

	//初始化文件
	config := Config{}

	yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatalf("Failed to unmarshal YAML: %v", err)
	}

	return config.Mysql
}
