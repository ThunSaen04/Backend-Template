package middleware

import (
	"strings"

	"backend-template/internal/modules/auth/utils"

	"github.com/gofiber/fiber/v3"
)

// AuthMiddleware validates the JWT token from the Authorization header
// and stores user claims in Fiber's Locals for downstream handlers
func AuthMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		// Get Authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"message": "Missing authorization header",
			})
		}

		// Extract token from "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"message": "Invalid authorization header format",
			})
		}

		tokenString := parts[1]

		// Parse and validate token
		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"message": "Invalid or expired token",
			})
		}

		// Store user data in Locals for downstream handlers
		c.Locals("user_id", claims.UserID)
		c.Locals("user_role", claims.Role)

		return c.Next()
	}
}
