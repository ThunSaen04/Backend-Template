package utils

// MessageResponseDTO represents a generic API response
type MessageResponseDTO struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"Operation successful"`
}

// DataResponseDTO represents an API response with data payload
type DataResponseDTO struct {
	Success bool        `json:"success" example:"true"`
	Message string      `json:"message" example:"Operation successful"`
	Data    interface{} `json:"data"`
}

// ErrorResponseDTO represents an error API response
type ErrorResponseDTO struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"An error occurred"`
}
