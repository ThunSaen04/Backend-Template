package models

import (
	"time"

	"gorm.io/gorm"
)

// RefreshToken represents a stored refresh token for JWT token rotation
type RefreshToken struct {
	gorm.Model
	UserID    uint      `gorm:"not null;index" json:"user_id"`
	Token     string    `gorm:"type:varchar(500);uniqueIndex;not null" json:"token"`
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`
	Revoked   bool      `gorm:"default:false;not null" json:"revoked"`

	// Relationship
	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"-"`
}
