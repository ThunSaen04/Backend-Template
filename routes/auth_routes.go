package routes

import (
	"backend-template/internal/database"
	handler_auth "backend-template/internal/modules/auth/handler"
	repository_auth "backend-template/internal/modules/auth/repository"
	service_auth "backend-template/internal/modules/auth/service"
	"backend-template/middleware"

	"github.com/gofiber/fiber/v3"
)

// SetupAuthRoutes configures all authentication-related routes under /api/v1/auth
func SetupAuthRoutes(v1 fiber.Router) {
	authRepo := repository_auth.NewAuthRepository(database.DB)
	authService := service_auth.NewAuthService(authRepo)
	authHandler := handler_auth.NewAuthHandler(authService)

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
