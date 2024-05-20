package handler

import (
	"context"
	"fmt"
	"login_module/internal/service"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

var ctx = context.Background()

type UserHandler struct {
	userService service.UserService
	redisClient *redis.Client
}

func NewUserHandler(userService service.UserService, redisClient *redis.Client) *UserHandler {
	return &UserHandler{userService: userService, redisClient: redisClient}
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	user := h.userService.GetUser(id)

	// Redis에 사용자 ID 저장
	err := h.redisClient.Set(ctx, fmt.Sprintf("user:%d", id), user, 0).Err()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to set user in Redis")
	}

	
	return c.SendString(user)
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	// 사용자 생성 로직
	return c.SendStatus(fiber.StatusCreated)
}
