package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("students.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	DB = database
}