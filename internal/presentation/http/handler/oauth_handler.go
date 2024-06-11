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
	providers    map[string]*oauth2.Config
}

func NewOAuthHandler(oAuthService *service.OAuthService, providers map[string]*oauth2.Config) *OAuthHandler {
	return &OAuthHandler{
		oauthService: oAuthService,
		providers:    providers,
	}
}

func (h *OAuthHandler) OAuthCallback(c *gin.Context) {
	m := dto.OAuthDTO{
		Code:     c.Query("code"),
		Provider: h.providers[c.Param("provider")],
	}
	res, err := h.oauthService.Login(c.Request.Context(), m)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie("id_token", res.IDToken, res.ExpiresIn, "/", "localhost", false, true)
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
	url := h.providers[provider].AuthCodeURL(oauthState, oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *OAuthHandler) Logout(c *gin.Context) {
	err := h.oauthService.Logout(c)
	if err != nil {
		log.Println("Failed to blacklist ID token:", err)
		return
	}
	c.SetCookie("id_token", "", -1, "/", "localhost", false, true)
	c.Redirect(http.StatusTemporaryRedirect, "http://localhost:5173/")
}
