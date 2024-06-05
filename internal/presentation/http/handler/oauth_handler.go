package handler

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"login_module/internal/application/dto"
	"login_module/internal/application/service"
	"login_module/internal/infrastructure/config"
	"login_module/pkg/util"
	"net/http"
	"time"

	"golang.org/x/oauth2"

	"github.com/gin-gonic/gin"
)

type OAuthHandler struct {
	OAuthService *service.OAuthService
}

func NewOAuthHandler(oAuthService *service.OAuthService) *OAuthHandler {
	return &OAuthHandler{
		OAuthService: oAuthService,
	}
}

func (h *OAuthHandler) OAuthCallback(c *gin.Context) {
	code := c.Query("code")
	user, err := getUserDataFromGoogle(code)
	if err != nil {
		log.Println(err.Error())
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	m := dto.OAuthDTO{
		UserUUID: user.UserInfo.ID,
		Provider: c.Param("provider"),
		Token: dto.AuthToken{
			IDToken:      user.IDToken,
			RefreshToken: user.RefreshToken,
			ExpiresIn:    user.ExpiresIn,
		},
	}
	if err := h.OAuthService.Login(c.Request.Context(), m); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	accessExp, _ := util.GetenvInt("HTTP_COOKIE_ACCESS_EXPIRY")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie("id_token", m.Token.IDToken, accessExp, "/", "localhost", false, true)
	c.Redirect(http.StatusTemporaryRedirect, "http://localhost:5173/")
}

func (h *OAuthHandler) BeginGoogleAuth(c *gin.Context) {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	oauthState := base64.URLEncoding.EncodeToString(b)
	u := config.GoogleOauthConfig.AuthCodeURL(oauthState, oauth2.AccessTypeOffline, oauth2.ApprovalForce)
	c.Redirect(http.StatusTemporaryRedirect, u)
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

func getUserDataFromGoogle(code string) (*dto.OAuthResponse, error) {
	token, err := config.GoogleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read response: %s", err.Error())
	}
	oauthRes := dto.OAuthResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		TokenType:    token.TokenType,
		Expiry:       token.Expiry,
	}
	if !token.Expiry.IsZero() {
		oauthRes.ExpiresIn = time.Until(token.Expiry).Seconds()
	}
	if idToken, ok := token.Extra("id_token").(string); ok {
		oauthRes.IDToken = idToken
	}
	if err := json.Unmarshal(contents, &oauthRes.UserInfo); err != nil {
		return nil, err
	}
	return &oauthRes, nil
}
