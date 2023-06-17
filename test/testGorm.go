package main

import (
	"fmt"
	"ginchat/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:010119@(localhost)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("链接数据库失败，err:%v\n", err)
		panic(err)
	}

	//db.AutoMigrate(&models.UserBasic{})
	db.AutoMigrate(&models.Community{})
	//db.AutoMigrate(&models.Message{})
	//db.AutoMigrate(&models.GroupBasic{})

	//user := models.Message{}
	//db.Create(&user)

}
