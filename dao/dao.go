package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dB *gorm.DB
)

func InitDB() {
	dsn := "root:lmh123@tcp(127.0.0.1:3306)/douban?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database")
	}

	dB = db
}
