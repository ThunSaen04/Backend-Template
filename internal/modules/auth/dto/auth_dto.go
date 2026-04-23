package dto_auth

// RegisterRequest represents the payload for user registration
type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email" example:"user@example.com"`
	Password string `json:"password" validate:"required,min=6" example:"secret123"`
}

// LoginRequest represents the payload for user login
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email" example:"user@example.com"`
	Password string `json:"password" validate:"required" example:"secret123"`
}

// RefreshRequest represents the payload for token refresh
type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required" example:"eyJhbGciOiJIUzI1NiIs..."`
}

// AuthResponse represents the response after successful authentication
type AuthResponse struct {
	AccessToken  string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIs..."`
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIs..."`
	Role         string `json:"role" example:"member"`
	UserID       uint   `json:"user_id" example:"1"`
}

// ProfileData represents user profile information
type ProfileData struct {
	ID    uint   `json:"id" example:"1"`
	Email string `json:"email" example:"user@example.com"`
	Role  string `json:"role" example:"member"`
}
