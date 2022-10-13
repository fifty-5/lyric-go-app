package utils

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection() {
	var error error
	DB, error = gorm.Open(mysql.Open(os.Getenv("DB_URL")), &gorm.Config{})

	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("Database connection successful")
	}
}
