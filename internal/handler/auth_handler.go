package handler

import (
	"login_module/internal/service"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// 로그인 요청 처리
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	accessToken, refreshToken, err := h.authService.Login(username, password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid credentials"})
	}

	// 리프레시 토큰을 쿠키로 설정
	c.Cookie(&fiber.Cookie{
		Name:     "refreshToken",
		Value:    refreshToken,
		HTTPOnly: true,
	})

	return c.JSON(fiber.Map{"accessToken": accessToken})
}

// 리프레시 토큰으로 새 액세스 토큰 발급
func (h *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	refreshToken := c.Cookies("refreshToken")

	newAccessToken, err := h.authService.RefreshToken(refreshToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid or expired refresh token"})
	}

	return c.JSON(fiber.Map{"accessToken": newAccessToken})
}

// 사용자 로그아웃 처리
func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	refreshToken := c.Cookies("refreshToken")

	if err := h.authService.Logout(refreshToken); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to logout"})
	}

	// 리프레시 토큰 쿠키 삭제
	c.ClearCookie("refreshToken")

	return c.SendStatus(fiber.StatusOK)
}
