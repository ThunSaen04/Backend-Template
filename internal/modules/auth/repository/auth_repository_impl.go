package repository

import (
	"backend-template/internal/models"

	"gorm.io/gorm"
)

// authRepositoryImpl is the concrete implementation of AuthRepository
type authRepositoryImpl struct {
	db *gorm.DB
}

// NewAuthRepository creates a new instance of AuthRepository
func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepositoryImpl{db: db}
}

// ── User Operations ─────────────────────────────────────────────────────────

// CreateUser inserts a new user record into the database
func (r *authRepositoryImpl) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

// FindByEmail retrieves a user by their email address
func (r *authRepositoryImpl) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByID retrieves a user by their ID
func (r *authRepositoryImpl) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// ── Refresh Token Operations ────────────────────────────────────────────────

// SaveRefreshToken stores a new refresh token in the database
func (r *authRepositoryImpl) SaveRefreshToken(token *models.RefreshToken) error {
	return r.db.Create(token).Error
}

// FindRefreshToken retrieves a refresh token by its token string
func (r *authRepositoryImpl) FindRefreshToken(tokenString string) (*models.RefreshToken, error) {
	var token models.RefreshToken
	err := r.db.Where("token = ? AND revoked = ?", tokenString, false).First(&token).Error
	if err != nil {
		return nil, err
	}
	return &token, nil
}

// DeleteRefreshToken soft-revokes a refresh token by marking it as revoked
func (r *authRepositoryImpl) DeleteRefreshToken(tokenString string) error {
	return r.db.Model(&models.RefreshToken{}).
		Where("token = ?", tokenString).
		Update("revoked", true).Error
}

// DeleteRefreshTokensByUserID revokes all refresh tokens for a given user
func (r *authRepositoryImpl) DeleteRefreshTokensByUserID(userID uint) error {
	return r.db.Model(&models.RefreshToken{}).
		Where("user_id = ?", userID).
		Update("revoked", true).Error
}
