package database

import (
	"fmt"
	"log"

	"backend-template/internal/config"
	"backend-template/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB is the global database instance
var DB *gorm.DB

// ConnectDB establishes a connection to PostgreSQL and runs AutoMigrate
func ConnectDB() {
	cfg := config.AppConfig

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Bangkok",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connected successfully")

	// AutoMigrate models
	err = DB.AutoMigrate(
		&models.User{},
		&models.RefreshToken{},
	)
	if err != nil {
		log.Fatalf("Failed to auto-migrate: %v", err)
	}

	log.Println("Database migration completed")
}
