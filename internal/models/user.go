package models

import "gorm.io/gorm"

// User represents the user entity in the database
type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Password string `gorm:"type:varchar(255);not null" json:"-"`
	Role     string `gorm:"type:varchar(50);default:'member';not null" json:"role"`
}
