package handler

import (
	"context"
	"log"
	"login_module/internal/application/service"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

type OAuthHandler struct {
	OAuthService *service.OAuthService
}

func NewOAuthHandler(OAuthService *service.OAuthService) *OAuthHandler {
	return &OAuthHandler{
		OAuthService: OAuthService,
	}
}

func InitOAuth() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	googleClientId := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET") // Replace with your SESSION_SECRET or similar
	maxAge := 86400 * 30                                    // 30 days
	isProd := false                                         // Set to true when serving over https

	store := sessions.NewCookieStore([]byte("test"))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = isProd

	gothic.Store = store

	googleProvider := google.New(googleClientId, googleClientSecret, "http://localhost:3000/api/oauth/google/callback")
	googleProvider.SetPrompt("consent")
	goth.UseProviders(googleProvider)
}

func (h *OAuthHandler) OAuthCallback(c *gin.Context) {
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	if err := h.OAuthService.StoreTokens(user.AccessToken, user.RefreshToken); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusAccepted, "login success")
}

func (h *OAuthHandler) BeginGoogleAuth(c *gin.Context) {
	provider := c.Param("provider")
	c.Request = c.Request.WithContext(context.WithValue(context.Background(), "provider", provider))
	q := c.Request.URL.Query()
	q.Add("provider", c.Param("provider"))
	c.Request.URL.RawQuery = q.Encode()
	gothic.BeginAuthHandler(c.Writer, c.Request)
}

func (h *OAuthHandler) Logout(c *gin.Context) {
	gothic.Logout(c.Writer, c.Request)
	c.Writer.Header().Set("Location", "/")
	c.Writer.WriteHeader(http.StatusTemporaryRedirect)
}
