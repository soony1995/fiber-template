package dto

import (
	"golang.org/x/oauth2"
	"time"
)

type OAuthDTO struct {
	Code     string         `json:"code"`
	Provider *oauth2.Config `json:"provider"`
}

type OAuthResponse struct {
	UserUUID              string    `json:"user_uuid"`
	TokenType             string    `json:"token_type,omitempty"`
	AccessToken           string    `json:"access_token"`
	RefreshToken          string    `json:"refresh_token,omitempty"`
	IDToken               string    `json:"id_token,omitempty"`
	ExpiresIn             int       `json:"expires_in,omitempty"`
	RefreshTokenExpiresIn int       `json:"refresh_token_expires_in,omitempty"`
	Expiry                time.Time `json:"expiry,omitempty"`
}
