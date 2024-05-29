package handler

import (
	"login_module/internal/application/dto"
	"login_module/internal/application/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// GetUser godoc
// @Summary Get user by ID
// @Description Get user details by ID
// @Tags user
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} dto.UserDTO
// @Failure 404 {object} map[string]string
// @Router /users/{id} [get]
func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	userDTO, err := h.userService.GetUser(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, userDTO)
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the given details
// @Tags user
// @Accept json
// @Produce json
// @Param user body dto.UserDTO true "User DTO"
// @Success 201 {object} dto.UserDTO
// @Failure 400 {object} map[string]string
// @Router /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	userDTO := new(dto.UserDTO)
	if err := c.ShouldBindJSON(userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot parse user data"})
		return
	}
	if err := h.userService.CreateUser(c, userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, userDTO)
}
