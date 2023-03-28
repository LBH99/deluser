package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	gorm.Model
	Id   int
	Name string `gorm:"varchar(30),"`
	//Password string `gorm:"varchar(50)"`
}

func Init() {
	dns := "root:123456@tcp(192.168.0.233:3306)/demoa?charset=utf8mb4&parseTime=True&loc=Local"
	open, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		return
	}

	DB = open

	//DB.AutoMigrate(&User{})
}
