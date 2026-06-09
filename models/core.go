package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

// https://gorm.io/zh_CN/docs/connecting_to_the_database.html

func init() {
	dsn := "root:root@tcp(127.0.0.1:3306)/go_study?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
}
