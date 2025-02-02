package config

import (
	"github.com/oliveirabalsa/2tp-management-backend/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("2tp-management.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	DB.Exec("PRAGMA journal_mode=WAL;")

	err = DB.AutoMigrate(&models.User{}, &models.Board{}, &models.Column{}, &models.Task{})
	if err != nil {
		log.Fatal("Failed to migrate database models: ", err)
	}

	log.Println("Database connected and migrated")
}
