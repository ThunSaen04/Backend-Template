package database

import (
	"backend-template/internal/config"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDBSQLite() {
	cfg := config.AppConfig

	var err error
	DB, err = gorm.Open(sqlite.Open(cfg.DBHost), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connected successfully")
}
