package config

import (
	"github.com/aceino/belajar_restapi/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connector between database
var DB *gorm.DB

func ConnectionDB() {
	dsn := "root:password@tcp(127.0.0.1:3306)/rest_fiber?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Book{}, &models.User{})

	DB = db

}
