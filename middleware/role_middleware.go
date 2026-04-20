package middleware

import (
	"backend-template/internal/modules/auth/utils"
	apputils "backend-template/internal/utils"

	"github.com/gofiber/fiber/v3"
)

// RoleMiddleware checks if the authenticated user has sufficient permission
// based on the role hierarchy. Must be used AFTER AuthMiddleware.
func RoleMiddleware(requiredRole string) fiber.Handler {
	return func(c fiber.Ctx) error {
		// Get user role from Locals (set by AuthMiddleware)
		userRole, ok := c.Locals("user_role").(string)
		if !ok || userRole == "" {
			return apputils.ErrorResponse(c, fiber.StatusUnauthorized, "Unauthorized: role not found in token")
		}

		// Check if user's role meets the required permission level
		if !utils.HasPermission(userRole, requiredRole) {
			return apputils.ErrorResponse(c, fiber.StatusForbidden, "Forbidden: insufficient permissions")
		}

		return c.Next()
	}
}

