package config

import (
    "Airbnb-room/models"
    "log"
    "os"

    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    dsn := os.Getenv("DB_URL")
    if dsn == "" {
        log.Fatalf("DB_URL is not set in environment variables")
    }

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }

    sqlDB, pingErr := db.DB()
    if pingErr != nil {
        log.Fatalf("Failed to ping the database: %v", pingErr)
    }

    if err := sqlDB.Ping(); err != nil {
        log.Fatalf("Database ping failed: %v", err)
    } else {
        log.Println("Successfully connected to the database")
    }

    if err := db.AutoMigrate(&models.Room{}); err != nil {
        log.Fatalf("AutoMigrate failed: %v", err)
    } else {
        log.Println("Database migration completed successfully")
    }

    DB = db
}
