package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func NewDB() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/go-zero?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	err = db.AutoMigrate(&Device{}, &User{}, &Product{})
	if err != nil {
		log.Println("DB ERROR :", err)
	}
	DB = db
}
