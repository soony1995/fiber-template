package handler

import (
	"login_module/internal/application/dto"
	"login_module/internal/application/service"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// GetUser godoc
// @Summary      Get user by ID
// @Description  Get user details by ID
// @Tags         user
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  dto.UserDTO
// @Failure      404  {object}  map[string]string
// @Router       /user/{id} [get]
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	userDTO, err := h.userService.GetUser(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
	}
	return c.JSON(userDTO)
}

// CreateUser godoc
// @Summary      Create a new user
// @Description  Create a new user with the given details
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        user  body      dto.UserDTO  true  "User DTO"
// @Success      201  {object}  dto.UserDTO
// @Failure      400  {object}  map[string]string
// @Router       /user [post]
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
