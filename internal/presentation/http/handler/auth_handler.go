package handler

import (
	"login_module/internal/application/service"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// Login handles the login request
// @Summary Login
// @Description Handles user login
// @Tags auth
// @Accept json
// @Produce json
// @Param username body string true "Username"
// @Param password body string true "Password"
// @Failure 400 {object} string"
// @Failure 404 {object} string"
// @Success 200 {object} string
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	accessToken, refreshToken, err := h.authService.Login(username, password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid credentials"})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "refreshToken",
		Value:    refreshToken,
		HTTPOnly: true,
	})

	return c.JSON(fiber.Map{"accessToken": accessToken})
}

// RefreshToken godoc
// @Summary      Refreshes the access token
// @Description  Refresh the access token using the refresh token
// @Tags         auth
// @Produce      json
// @Success      200  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Router       /auth/refresh_token [post]

func (h *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	refreshToken := c.Cookies("refreshToken")

	newAccessToken, err := h.authService.RefreshToken(refreshToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid or expired refresh token"})
	}

	return c.JSON(fiber.Map{"accessToken": newAccessToken})
}

// Logout godoc
// @Summary      Logs out a user
// @Description  Logs out the user by deleting the refresh token
// @Tags         auth
// @Produce      json
// @Success      200
// @Failure      500  {object}  map[string]string
// @Router       /auth/logout [post]
func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	refreshToken := c.Cookies("refreshToken")

	if err := h.authService.Logout(refreshToken); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to logout"})
	}

	c.ClearCookie("refreshToken")

	return c.SendStatus(fiber.StatusOK)
}
