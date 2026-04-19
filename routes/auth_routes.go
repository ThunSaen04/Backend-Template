package routes

import (
	"backend-template/internal/database"
	"backend-template/internal/modules/auth/handler"
	"backend-template/internal/modules/auth/repository"
	"backend-template/internal/modules/auth/service"
	"backend-template/middleware"

	"github.com/gofiber/fiber/v3"
)

// SetupAuthRoutes configures all authentication-related routes under /api/v1/auth
func SetupAuthRoutes(v1 fiber.Router) {
	authRepo := repository.NewAuthRepository(database.DB)
	authService := service.NewAuthService(authRepo)
	authHandler := handler.NewAuthHandler(authService)

	auth := v1.Group("/auth")

	// Public routes (no authentication required)
	auth.Post("/register", authHandler.Register, middleware.StrictAuthLimiter())
	auth.Post("/login", authHandler.Login, middleware.StrictAuthLimiter())
	auth.Post("/refresh", authHandler.RefreshToken)
	auth.Get("/health", authHandler.HealthCheck)

	// Protected routes (authentication required)
	protected := auth.Group("", middleware.AuthMiddleware())
	protected.Get("/profile", authHandler.GetProfile)
	protected.Post("/logout", authHandler.Logout)

	// Admin-only routes (authentication + admin role required)
	admin := auth.Group("", middleware.AuthMiddleware(), middleware.RoleMiddleware("admin"))
	admin.Get("/users", authHandler.GetAllUsers)
}
