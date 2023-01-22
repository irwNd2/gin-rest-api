package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:mariadb@tcp(localhost:3306)/gin_restapi"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Product{})

	DB = database
}
