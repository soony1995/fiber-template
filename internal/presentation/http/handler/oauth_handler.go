package handler

import (
	"crypto/rand"
	"encoding/base64"
	"golang.org/x/oauth2"
	"log"
	"login_module/internal/application/dto"
	"login_module/internal/application/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OAuthHandler struct {
	oauthService *service.OAuthService
	config       map[string]*oauth2.Config
}

func NewOAuthHandler(oAuthService *service.OAuthService, providers map[string]*oauth2.Config) *OAuthHandler {
	return &OAuthHandler{
		oauthService: oAuthService,
		config:       providers,
	}
}

func (h *OAuthHandler) OAuthCallback(c *gin.Context) {
	provider := c.Param("provider")
	m := dto.OAuthDTO{
		Code:     c.Query("code"),
		Provider: provider,
		Config:   h.config[provider],
	}
	res, err := h.oauthService.Login(c.Request.Context(), m)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie(provider, res.IDToken, res.TokenExpiresIn, "/", "localhost", false, true)
	c.Redirect(http.StatusTemporaryRedirect, "http://localhost:5173/")
}

func (h *OAuthHandler) BeginOAuth(c *gin.Context) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	oauthState := base64.URLEncoding.EncodeToString(b)
	provider := c.Param("provider")
	// TODO: kakao의 경우 동의 화면이 나타나지 않고 바로 로그인되는 현상 발견
	url := h.config[provider].AuthCodeURL(oauthState, oauth2.AccessTypeOffline, oauth2.ApprovalForce)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *OAuthHandler) Logout(c *gin.Context) {
	err := h.oauthService.Logout(c)
	if err != nil {
		log.Println("Failed to blacklist ID token:", err)
		return
	}
	c.SetCookie(c.Param("provider"), "", -1, "/", "localhost", false, true)
	c.Redirect(http.StatusTemporaryRedirect, "http://localhost:5173/")
}
