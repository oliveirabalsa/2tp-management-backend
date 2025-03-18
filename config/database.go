package config

import (
	"log"

	"github.com/google/uuid"
	"github.com/oliveirabalsa/2tp-management-backend/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("2tp-management.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Register callback for UUID generation
	DB.Callback().Create().Before("gorm:create").Register("generate_uuid", func(db *gorm.DB) {
		if model, ok := db.Statement.Dest.(interface{ GetID() uuid.UUID }); ok {
			if model.GetID() == uuid.Nil {
				db.Statement.SetColumn("ID", uuid.New())
			}
		}
	})

	DB.Exec("PRAGMA journal_mode=WAL;")

	err = DB.AutoMigrate(&models.User{}, &models.Board{}, &models.Column{}, &models.Task{})
	if err != nil {
		log.Fatal("Failed to migrate database models: ", err)
	}

	log.Println("Database connected and migrated")
}
