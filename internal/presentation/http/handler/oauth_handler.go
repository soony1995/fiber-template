package handler

import (
	"crypto/rand"
	"encoding/base64"
	"golang.org/x/oauth2"
	"log"
	"login_module/internal/application/dto"
	"login_module/internal/application/service/oauth"
	"login_module/internal/infrastructure/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OAuthHandler struct {
	OAuthService *oauth.OAuthService
}

func NewOAuthHandler(oAuthService *oauth.OAuthService) *OAuthHandler {
	return &OAuthHandler{
		OAuthService: oAuthService,
	}
}

func (h *OAuthHandler) OAuthCallback(c *gin.Context) {
	m := dto.OAuthDTO{
		Code:     c.Query("code"),
		Provider: c.Param("provider"),
	}
	res, err := h.OAuthService.Login(c.Request.Context(), m)
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
	var url string
	switch provider {
	case "google":
		url = config.GoogleOauthConfig.AuthCodeURL(oauthState, oauth2.AccessTypeOffline, oauth2.ApprovalForce)
	case "kakao":
		url = config.KakaoOauthConfig.AuthCodeURL(oauthState, oauth2.AccessTypeOffline)
	}
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *OAuthHandler) Logout(c *gin.Context) {
	err := h.OAuthService.Logout(c)
	if err != nil {
		// Redis 저장 실패 시 로그 기록 (선택 사항)
		log.Println("Failed to blacklist ID token:", err)
		return
	}
	c.SetCookie("id_token", "", -1, "/", "localhost", false, true)
	c.Redirect(http.StatusTemporaryRedirect, "http://localhost:5173/")
}

/*
로그인 시작 시: 사용자가 로그인 버튼을 클릭하면, 서버는 고유한 상태 값을 생성하고 이 값을 oauthstate 쿠키에 저장합니다.
그 후, 이 상태 값을 포함하여 OAuth 제공자(예: Google)에 로그인 요청을 보냅니다.

OAuth 콜백 처리: 사용자가 OAuth 제공자에서 인증을 마치면, 제공자는 사용자를 리디렉션 URI로 돌려보내면서 초기에 보낸 상태 값(state)을 반환합니다.
서버는 반환된 상태 값과 oauthstate 쿠키에 저장된 값을 비교합니다. 이 값들이 일치하면 요청이 유효한 것으로 간주하고, 그렇지 않으면 요청을 거부합니다.
*/
