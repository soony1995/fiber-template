package handler

import (
	"github.com/gin-contrib/sessions"
	"login_module/internal/application/dto"
	"login_module/internal/application/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(c *gin.Context) {
	// session은 handler에서 처리하는 것이 맞고, redis에 저장하는 로직은 서비스 레이어에서 하는 것이 맞으나,
	// redis에 저장하려는 값이 service 로직이 완료된 후에야 알 수 있어 session ID를 알지 못해 저장하기 어렵다.
	// redis에 저장하는 로직을 handler 레이어로 빼내어 session ID를 알게 된 후에 저장하는 방법을 사용할 수 있다.
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	user, err := h.authService.Login(c, req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	session := sessions.Default(c)
	session.Set("username", user.Username)
	session.Options(sessions.Options{MaxAge: 600}) // 10 minutes
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "login successful"})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, gin.H{"message": "logout successful"})
}
