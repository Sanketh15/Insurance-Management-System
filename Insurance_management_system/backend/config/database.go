package config

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"backend/models" // Import the models package where your model structs are defined
)

var DB *gorm.DB

// ConnectDatabase initializes the database connection and performs auto-migration
func ConnectDatabase() {
	// Get the database connection string from environment variable
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatalf("DATABASE_URL is not set in the environment variables")
	}

	// Open a connection to the PostgreSQL database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Perform auto-migration to create or update tables based on the models
	err = db.AutoMigrate(&models.Application{}, &models.Feedback{},&models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate the database: %v", err)
	}

	// Set the global DB variable for use throughout the application
	DB = db
}
// GetDB returns the database instance
func GetDB() *gorm.DB {
	if DB == nil {
		log.Fatal("Database connection is not initialized. Call ConnectDatabase first.")
	}
	return DB
}