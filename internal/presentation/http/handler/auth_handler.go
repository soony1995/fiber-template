package handler

import (
	"login_module/internal/application/service/oauth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *oauth.AuthService
}

func NewAuthHandler(authService *oauth.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// Login godoc
// @Summary Login
// @Description Handles user login
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	accessToken, refreshToken, err := h.authService.Login(credentials.Username, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	c.SetCookie("refreshToken", refreshToken, 3600*24, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"accessToken": accessToken})
}

// RefreshToken godoc
// @Summary Refreshes the access token
// @Description Refresh the access token using the refresh token
// @Tags auth
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /auth/refresh_token [post]
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	refreshToken, err := c.Cookie("refreshToken")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired refresh token"})
		return
	}

	newAccessToken, err := h.authService.RefreshToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired refresh token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"accessToken": newAccessToken})
}

// Logout godoc
// @Summary Logs out a user
// @Description Logs out the user by deleting the refresh token
// @Tags auth
// @Produce json
// @Success 200
// @Failure 500 {object} map[string]string
// @Router /auth/logout [post]
func (h *AuthHandler) Logout(c *gin.Context) {
	refreshToken, err := c.Cookie("refreshToken")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to logout"})
		return
	}

	if err := h.authService.Logout(refreshToken); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to logout"})
		return
	}

	c.SetCookie("refreshToken", "", -1, "/", "", false, true)
	c.Status(http.StatusOK)
}
