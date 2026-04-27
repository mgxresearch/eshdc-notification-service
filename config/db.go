package config

import (
	"fmt"
	"log"
	"os"

	"github.com/eshdc/notification-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		// Default for local development
		dsn = "host=localhost user=postgres password=postgres dbname=eshdc_notification port=5432 sslmode=disable"
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	fmt.Println("Database connection established")

	// Migrate models
	err = DB.AutoMigrate(&models.NotificationTemplate{}, &models.Notification{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}
	
	fmt.Println("Database migration completed")
}
