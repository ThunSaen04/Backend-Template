package service_auth

import (
	"backend-template/internal/models"
	dto_auth "backend-template/internal/modules/auth/dto"
)

// AuthService defines the interface for authentication business logic
type AuthService interface {
	Register(req *dto_auth.RegisterRequest) (*dto_auth.AuthResponse, error)
	Login(req *dto_auth.LoginRequest) (*dto_auth.AuthResponse, error)
	RefreshToken(req *dto_auth.RefreshRequest) (*dto_auth.AuthResponse, error)
	Logout(refreshToken string) error
	GetProfile(userID uint) (*models.User, error)
}
