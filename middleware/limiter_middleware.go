package middleware

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/limiter"
)

// APILimiter defines a general rate limit for all API endpoints.
// Example: Max 100 requests per 1 minute per IP.
func APILimiter() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        100,
		Expiration: 1 * time.Minute,
		LimitReached: func(c fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"success": false,
				"message": "Too many requests. Please try again later.",
			})
		},
	})
}

// StrictAuthLimiter defines a stricter rate limit for sensitive endpoints like Login/Register
// to prevent brute-force and credential stuffing attacks.
// Example: Max 5 requests per 1 minute per IP.
func StrictAuthLimiter() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        5,
		Expiration: 1 * time.Minute,
		LimitReached: func(c fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"success": false,
				"message": "Too many attempts. Please wait before trying again.",
			})
		},
	})
}
