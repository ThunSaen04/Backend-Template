package service_auth

import (
	"errors"
	"time"

	"backend-template/internal/config"
	"backend-template/internal/models"
	dto_auth "backend-template/internal/modules/auth/dto"
	repository_auth "backend-template/internal/modules/auth/repository"
	utils_auth "backend-template/internal/modules/auth/utils"

	"gorm.io/gorm"
)

// AuthService defines the interface for authentication business logic
type AuthService interface {
	Register(req *dto_auth.RegisterRequest) (*dto_auth.AuthResponse, error)
	Login(req *dto_auth.LoginRequest) (*dto_auth.AuthResponse, error)
	RefreshToken(req *dto_auth.RefreshRequest) (*dto_auth.AuthResponse, error)
	Logout(refreshToken string) error
	GetProfile(userID uint) (*models.User, error)
}

// authServiceImpl is the concrete implementation of AuthService
type authServiceImpl struct {
	repo repository_auth.AuthRepository
}

// NewAuthService creates a new instance of AuthService
func NewAuthService(repo repository_auth.AuthRepository) AuthService {
	return &authServiceImpl{repo: repo}
}

// Register handles user registration with password hashing and default role assignment
func (s *authServiceImpl) Register(req *dto_auth.RegisterRequest) (*dto_auth.AuthResponse, error) {
	// Check if user already exists
	existingUser, err := s.repo.FindByEmail(req.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("failed to check existing user")
	}
	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	// Hash the password
	hashedPassword, err := utils_auth.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// Create user with default role "member"
	user := &models.User{
		Email:    req.Email,
		Password: hashedPassword,
		Role:     "member",
	}

	if err := s.repo.CreateUser(user); err != nil {
		return nil, errors.New("failed to create user")
	}

	// Generate tokens
	return s.generateTokenPair(user)
}

// Login validates credentials and returns access + refresh tokens
func (s *authServiceImpl) Login(req *dto_auth.LoginRequest) (*dto_auth.AuthResponse, error) {
	// Find user by email
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Compare password
	if !utils_auth.ComparePassword(user.Password, req.Password) {
		return nil, errors.New("invalid email or password")
	}

	// Generate tokens
	return s.generateTokenPair(user)
}

// RefreshToken validates a refresh token and generates a new access token
func (s *authServiceImpl) RefreshToken(req *dto_auth.RefreshRequest) (*dto_auth.AuthResponse, error) {
	// Parse the refresh token to get claims
	claims, err := utils_auth.ParseToken(req.RefreshToken)
	if err != nil {
		return nil, errors.New("invalid or expired refresh token")
	}

	// Verify it's actually a refresh token
	if claims.Subject != "refresh" {
		return nil, errors.New("invalid token type")
	}

	// Check if token exists in database and is not revoked
	storedToken, err := s.repo.FindRefreshToken(req.RefreshToken)
	if err != nil {
		return nil, errors.New("refresh token not found or revoked")
	}

	// Check if token is expired
	if storedToken.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("refresh token expired")
	}

	// Find the user
	user, err := s.repo.FindByID(claims.UserID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Revoke the old refresh token (token rotation)
	_ = s.repo.DeleteRefreshToken(req.RefreshToken)

	// Generate new token pair
	return s.generateTokenPair(user)
}

// Logout revokes the given refresh token
func (s *authServiceImpl) Logout(refreshToken string) error {
	err := s.repo.DeleteRefreshToken(refreshToken)
	if err != nil {
		return errors.New("failed to revoke token")
	}
	return nil
}

// GetProfile retrieves user profile by ID
func (s *authServiceImpl) GetProfile(userID uint) (*models.User, error) {
	user, err := s.repo.FindByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// generateTokenPair creates an access token and a refresh token, saving the refresh token to DB
func (s *authServiceImpl) generateTokenPair(user *models.User) (*dto_auth.AuthResponse, error) {
	// Generate access token (6h)
	accessToken, err := utils_auth.GenerateAccessToken(user.ID, user.Role)
	if err != nil {
		return nil, errors.New("failed to generate access token")
	}

	// Generate refresh token (24h)
	refreshToken, err := utils_auth.GenerateRefreshToken(user.ID, user.Role)
	if err != nil {
		return nil, errors.New("failed to generate refresh token")
	}

	// Store refresh token in database
	storedToken := &models.RefreshToken{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(config.AppConfig.RefreshTokenDuration),
		Revoked:   false,
	}

	if err := s.repo.SaveRefreshToken(storedToken); err != nil {
		return nil, errors.New("failed to save refresh token")
	}

	return &dto_auth.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Role:         user.Role,
		UserID:       user.ID,
	}, nil
}
