package database

import (
	"backend-template/internal/models"
	"log"
)

func AutoMigrate() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.RefreshToken{},
	)
	if err != nil {
		log.Fatalf("Failed to auto-migrate: %v", err)
	}
}
