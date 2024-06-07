package oauth

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"login_module/internal/application/dto"
	"login_module/internal/domain/model"
	"login_module/internal/domain/repository"
	"login_module/internal/infrastructure/config"
	"login_module/pkg/jwt"
)

type OAuthService struct {
	redisClient repository.AuthRepository
}

func NewOAuthService(r repository.AuthRepository) *OAuthService {
	return &OAuthService{
		redisClient: r,
	}
}

func (s *OAuthService) Login(c context.Context, req dto.OAuthDTO) (res *dto.OAuthResponse, err error) {
	res, err = processOAuthToken(c, req)
	if err != nil {
		return nil, err
	}

	// TODO
	//check m.UserUuid in mysql server
	//if not exist redirect register page

	SaveRefreshToken := model.SaveRefreshToken{
		UserUUID:              res.UserUUID,
		RefreshToken:          res.RefreshToken,
		RefreshTokenExpiresIn: res.RefreshTokenExpiresIn,
	}
	if err := s.redisClient.SaveRefreshToken(c, SaveRefreshToken); err != nil {
		return nil, err
	}
	return res, nil
}

func processOAuthToken(c context.Context, req dto.OAuthDTO) (res *dto.OAuthResponse, err error) {
	var token *oauth2.Token
	res = &dto.OAuthResponse{}
	switch req.Provider {
	case "google":
		token, err = config.GoogleOauthConfig.Exchange(c, req.Code)
		if err != nil {
			return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
		}
	case "kakao":
		token, err = config.KakaoOauthConfig.Exchange(c, req.Code)
		if err != nil {
			return nil, fmt.Errorf("failed to exchange code: %s", err.Error())
		}
	}
	res.AccessToken = token.AccessToken
	res.TokenType = token.TokenType
	res.RefreshToken = token.RefreshToken
	res.Expiry = token.Expiry
	if ExpiresIn, ok := token.Extra("expires_in").(int); ok {
		res.ExpiresIn = ExpiresIn
	}
	if RefreshTokenExpiresIn, ok := token.Extra("refresh_token_expires_in").(int); ok {
		res.RefreshTokenExpiresIn = RefreshTokenExpiresIn
	}
	if idToken, ok := token.Extra("id_token").(string); ok {
		res.IDToken = idToken
	}
	id, err := jwt.ParseIDToken(res.IDToken)
	if err != nil {
		return nil, fmt.Errorf("failed to parse id token: %s", err.Error())
	}
	res.UserUUID = id
	return res, nil
}

func (s *OAuthService) Logout(ctx context.Context) error {
	userUUID, ok := ctx.Value("userUUID").(string) // 문자열 타입으로 변환 (타입 단언)
	if !ok {
		return fmt.Errorf("userUUID not found in context or not a string")
	}
	err := s.redisClient.DeleteIDToken(ctx, userUUID)
	if err != nil {
		return err
	}
	return nil
}
