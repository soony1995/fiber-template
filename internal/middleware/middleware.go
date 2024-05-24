package middleware

import (
	"log"
	"login_module/internal/pkg/jwt"

	"github.com/gofiber/fiber/v2"
)

// TokenValid 미들웨어는 요청에서 JWT 액세스 토큰의 유효성을 검사합니다.
func TokenValid(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "no authentication token provided"})
	}

	// "Bearer " 접두사 제거
	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}

	claims, err := jwt.ParseToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid or expired token"})
	}
	log.Print("여기 들어옴.")

	// 클레임을 컨텍스트에 추가
	c.Locals("username", claims.Username)
	return c.Next()
}
