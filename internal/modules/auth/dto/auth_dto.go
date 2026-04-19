package dto

// RegisterRequest represents the payload for user registration
type RegisterRequest struct {
	Email    string `json:"email" example:"user@example.com"`
	Password string `json:"password" example:"secret123"`
}

// LoginRequest represents the payload for user login
type LoginRequest struct {
	Email    string `json:"email" example:"user@example.com"`
	Password string `json:"password" example:"secret123"`
}

// RefreshRequest represents the payload for token refresh
type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIs..."`
}

// AuthResponse represents the response after successful authentication
type AuthResponse struct {
	AccessToken  string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIs..."`
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIs..."`
	Role         string `json:"role" example:"member"`
	UserID       uint   `json:"user_id" example:"1"`
}

// MessageResponse represents a generic API response
type MessageResponse struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"Operation successful"`
}

// DataResponse represents an API response with data payload
type DataResponse struct {
	Success bool        `json:"success" example:"true"`
	Message string      `json:"message" example:"Operation successful"`
	Data    interface{} `json:"data"`
}

// ErrorResponse represents an error API response
type ErrorResponse struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"An error occurred"`
}

// ProfileData represents user profile information
type ProfileData struct {
	ID    uint   `json:"id" example:"1"`
	Email string `json:"email" example:"user@example.com"`
	Role  string `json:"role" example:"member"`
}
