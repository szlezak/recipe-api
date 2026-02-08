package database

import (
	"os"

	"github.com/szlezak/recipe-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DATABASE_URL")

	// If it's empty (local dev), you can provide a fallback or panic
	if dsn == "" {
		panic("DATABASE_URL environment variable is not set")
	}

	dbPath := os.Getenv("DB_PATH")
    if dbPath == "" {
        dbPath = "recipes.db"
    }

	

	database, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&models.Recipe{})
	DB = database
}