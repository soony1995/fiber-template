package config

import (
	"golang.org/x/oauth2/kakao"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var GoogleOauthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:3000/api/oauth/google/callback",
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/photoslibrary.readonly"},
	Endpoint:     google.Endpoint,
}

var KakaoOauthConfig = &oauth2.Config{
	RedirectURL: "http://localhost:3000/api/oauth/kakao/callback",
	ClientID:    os.Getenv("KAKAO_CLIENT_ID"),
	Endpoint:    kakao.Endpoint,
}
