package utils

import "github.com/gofiber/fiber/v3"

// Response is the standard API response structure used across all endpoints.
// All fields except "success" use omitempty to keep responses clean.
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

// SuccessResponse sends a standardized success JSON response.
//
// Example:
//
//	return utils.SuccessResponse(c, fiber.StatusOK, "User fetched", user)
func SuccessResponse(c fiber.Ctx, statusCode int, message string, data interface{}) error {
	return c.Status(statusCode).JSON(Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// MessageResponse sends a standardized success JSON response without a data payload.
//
// Example:
//
//	return utils.MessageResponse(c, fiber.StatusOK, "Logged out successfully")
func MessageResponse(c fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(Response{
		Success: true,
		Message: message,
	})
}

// ErrorResponse sends a standardized error JSON response.
//
// Example:
//
//	return utils.ErrorResponse(c, fiber.StatusNotFound, "User not found")
func ErrorResponse(c fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(Response{
		Success: false,
		Message: message,
	})
}

// ValidationErrorResponse sends a 422 Unprocessable Entity response
// with detailed field-level validation errors.
//
// Example:
//
//	if errs := utils.ValidateStruct(&req); errs != nil {
//	    return utils.ValidationErrorResponse(c, errs)
//	}
func ValidationErrorResponse(c fiber.Ctx, errors []*ValidationError) error {
	return c.Status(fiber.StatusUnprocessableEntity).JSON(Response{
		Success: false,
		Message: "Validation failed",
		Errors:  errors,
	})
}

// PaginatedResponse sends a standardized success JSON response
// that includes pagination metadata in the "meta" field.
//
// Example:
//
//	return utils.PaginatedResponse(c, "Items fetched", items, meta)
func PaginatedResponse(c fiber.Ctx, message string, data interface{}, meta *PaginationMeta) error {
	return c.Status(fiber.StatusOK).JSON(Response{
		Success: true,
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}
