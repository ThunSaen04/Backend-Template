package handler

import (
	"backend-template/internal/modules/auth/dto"
	"backend-template/internal/modules/auth/service"
	"backend-template/internal/modules/auth/utils"
	apputils "backend-template/internal/utils"

	"github.com/gofiber/fiber/v3"
)

// AuthHandler handles HTTP requests for authentication
type AuthHandler struct {
	service service.AuthService
}

// NewAuthHandler creates a new instance of AuthHandler
func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

// Register godoc
// @Summary      Register a new user
// @Description  Create a new user account with email and password. Default role is "member".
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request  body      dto.RegisterRequest  true  "Registration data"
// @Success      201      {object}  dto.DataResponse{data=dto.AuthResponse}
// @Failure      400      {object}  dto.ErrorResponse
// @Failure      409      {object}  dto.ErrorResponse
// @Failure      422      {object}  dto.ErrorResponse
// @Router       /api/v1/auth/register [post]
func (h *AuthHandler) Register(c fiber.Ctx) error {
	var req dto.RegisterRequest

	if err := c.Bind().JSON(&req); err != nil {
		return apputils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	// Struct validation (automatic via validate tags)
	if errs := apputils.ValidateStruct(&req); errs != nil {
		return apputils.ValidationErrorResponse(c, errs)
	}

	response, err := h.service.Register(&req)
	if err != nil {
		return apputils.ErrorResponse(c, fiber.StatusConflict, err.Error())
	}

	return apputils.SuccessResponse(c, fiber.StatusCreated, "User registered successfully", response)
}

// Login godoc
// @Summary      Login user
// @Description  Authenticate user with email and password. Returns access token (6h) and refresh token (24h).
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request  body      dto.LoginRequest  true  "Login credentials"
// @Success      200      {object}  dto.DataResponse{data=dto.AuthResponse}
// @Failure      400      {object}  dto.ErrorResponse
// @Failure      401      {object}  dto.ErrorResponse
// @Failure      422      {object}  dto.ErrorResponse
// @Router       /api/v1/auth/login [post]
func (h *AuthHandler) Login(c fiber.Ctx) error {
	var req dto.LoginRequest

	if err := c.Bind().JSON(&req); err != nil {
		return apputils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	// Struct validation (automatic via validate tags)
	if errs := apputils.ValidateStruct(&req); errs != nil {
		return apputils.ValidationErrorResponse(c, errs)
	}

	response, err := h.service.Login(&req)
	if err != nil {
		return apputils.ErrorResponse(c, fiber.StatusUnauthorized, err.Error())
	}

	return apputils.SuccessResponse(c, fiber.StatusOK, "Login successful", response)
}

// RefreshToken godoc
// @Summary      Refresh access token
// @Description  Generate a new access token using a valid refresh token. Implements token rotation (old refresh token is revoked).
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request  body      dto.RefreshRequest  true  "Refresh token"
// @Success      200      {object}  dto.DataResponse{data=dto.AuthResponse}
// @Failure      400      {object}  dto.ErrorResponse
// @Failure      401      {object}  dto.ErrorResponse
// @Failure      422      {object}  dto.ErrorResponse
// @Router       /api/v1/auth/refresh [post]
func (h *AuthHandler) RefreshToken(c fiber.Ctx) error {
	var req dto.RefreshRequest

	if err := c.Bind().JSON(&req); err != nil {
		return apputils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	// Struct validation (automatic via validate tags)
	if errs := apputils.ValidateStruct(&req); errs != nil {
		return apputils.ValidationErrorResponse(c, errs)
	}

	response, err := h.service.RefreshToken(&req)
	if err != nil {
		return apputils.ErrorResponse(c, fiber.StatusUnauthorized, err.Error())
	}

	return apputils.SuccessResponse(c, fiber.StatusOK, "Token refreshed successfully", response)
}

// Logout godoc
// @Summary      Logout user
// @Description  Revoke the refresh token to invalidate the session.
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request  body      dto.RefreshRequest  true  "Refresh token to revoke"
// @Success      200      {object}  dto.MessageResponse
// @Failure      400      {object}  dto.ErrorResponse
// @Failure      401      {object}  dto.ErrorResponse
// @Failure      422      {object}  dto.ErrorResponse
// @Router       /api/v1/auth/logout [post]
func (h *AuthHandler) Logout(c fiber.Ctx) error {
	var req dto.RefreshRequest

	if err := c.Bind().JSON(&req); err != nil {
		return apputils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	// Struct validation (automatic via validate tags)
	if errs := apputils.ValidateStruct(&req); errs != nil {
		return apputils.ValidationErrorResponse(c, errs)
	}

	if err := h.service.Logout(req.RefreshToken); err != nil {
		return apputils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return apputils.MessageResponse(c, fiber.StatusOK, "Logged out successfully")
}

// GetProfile godoc
// @Summary      Get user profile
// @Description  Retrieve the authenticated user's profile information.
// @Tags         Auth
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  dto.DataResponse{data=dto.ProfileData}
// @Failure      401  {object}  dto.ErrorResponse
// @Failure      404  {object}  dto.ErrorResponse
// @Router       /api/v1/auth/profile [get]
func (h *AuthHandler) GetProfile(c fiber.Ctx) error {
	// Extract user_id from JWT claims stored in Locals by auth middleware
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return apputils.ErrorResponse(c, fiber.StatusUnauthorized, "Unauthorized")
	}

	user, err := h.service.GetProfile(userID)
	if err != nil {
		return apputils.ErrorResponse(c, fiber.StatusNotFound, err.Error())
	}

	return apputils.SuccessResponse(c, fiber.StatusOK, "Profile fetched successfully", fiber.Map{
		"id":    user.ID,
		"email": user.Email,
		"role":  user.Role,
	})
}

// GetAllUsers godoc
// @Summary      Get all users (Admin only)
// @Description  Retrieve a list of all users. Requires admin role.
// @Tags         Auth
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  dto.MessageResponse
// @Failure      401  {object}  dto.ErrorResponse
// @Failure      403  {object}  dto.ErrorResponse
// @Router       /api/v1/auth/users [get]
func (h *AuthHandler) GetAllUsers(c fiber.Ctx) error {
	// This endpoint is protected by role middleware (admin only)
	_ = c.Locals("user_id")

	return apputils.MessageResponse(c, fiber.StatusOK, "Admin access granted - user list endpoint")
}

// HealthCheck godoc
// @Summary      Health check
// @Description  Check if the auth service is running and view role hierarchy.
// @Tags         Auth
// @Produce      json
// @Success      200  {object}  dto.MessageResponse
// @Router       /api/v1/auth/health [get]
func (h *AuthHandler) HealthCheck(c fiber.Ctx) error {
	return apputils.SuccessResponse(c, fiber.StatusOK, "Auth service is healthy", fiber.Map{
		"version": "1.0.0",
		"roles":   utils.RoleHierarchy,
	})
}
