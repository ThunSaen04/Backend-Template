package repository_auth

import (
	"backend-template/internal/models"
)

// AuthRepository defines the interface for user and token data access operations
type AuthRepository interface {
	// User operations
	CreateUser(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	FindByID(id uint) (*models.User, error)

	// Refresh token operations
	SaveRefreshToken(token *models.RefreshToken) error
	FindRefreshToken(tokenString string) (*models.RefreshToken, error)
	DeleteRefreshToken(tokenString string) error
	DeleteRefreshTokensByUserID(userID uint) error
}
