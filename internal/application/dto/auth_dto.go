package dto

import (
	"golang.org/x/oauth2"
)

type OAuthDTO struct {
	Code     string         `json:"code"`
	Provider string         `json:"provider"`
	Config   *oauth2.Config `json:"config"`
}

type OAuthResponse struct {
	UserUUID              string `json:"user_uuid"`
	TokenType             string `json:"token_type,omitempty"`
	AccessToken           string `json:"access_token"`
	RefreshToken          string `json:"refresh_token,omitempty"`
	RefreshTokenExpiresIn int    `json:"refresh_token_expires_in,omitempty"`
	IDToken               string `json:"id_token,omitempty"`
	TokenExpiresIn        int    `json:"token_expires_in,omitempty"`
}
