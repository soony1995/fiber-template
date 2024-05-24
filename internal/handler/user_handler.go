package handler

import (
	"login_module/internal/dto"
	"login_module/internal/service"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// GetUser 사용자 정보 조회 핸들러
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	userDTO, err := h.userService.GetUser(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
	}
	return c.JSON(userDTO)
}

// CreateUser 새로운 사용자 생성 핸들러
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	userDTO := new(dto.UserDTO)
	if err := c.BodyParser(userDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse user data"})
	}
	if err := h.userService.CreateUser(c.Context(), userDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(userDTO)
}
