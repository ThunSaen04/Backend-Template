package middleware

import (
	"backend-template/internal/modules/auth/utils"

	"github.com/gofiber/fiber/v3"
)

// RoleMiddleware checks if the authenticated user has sufficient permission
// based on the role hierarchy. Must be used AFTER AuthMiddleware.
func RoleMiddleware(requiredRole string) fiber.Handler {
	return func(c fiber.Ctx) error {
		// Get user role from Locals (set by AuthMiddleware)
		userRole, ok := c.Locals("user_role").(string)
		if !ok || userRole == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"message": "Unauthorized: role not found in token",
			})
		}

		// Check if user's role meets the required permission level
		if !utils.HasPermission(userRole, requiredRole) {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"success": false,
				"message": "Forbidden: insufficient permissions",
			})
		}

		return c.Next()
	}
}
