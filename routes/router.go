package routes

import (
	"backend-template/middleware"

	"github.com/gofiber/fiber/v3"
)

// SetupRoutes initializes all route groups, wires dependencies, and configures API versioning
func SetupRoutes(app *fiber.App) {
	// API v1 group with general rate limit
	v1 := app.Group("/api/v1", middleware.APILimiter())

	// Setup module routes under v1
	SetupAuthRoutes(v1)

	// Root health check
	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"message": "Backend Template API v1.0.0",
			"docs":    "/swagger/index.html",
		})
	})
}
