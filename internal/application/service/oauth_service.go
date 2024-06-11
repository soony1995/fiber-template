package service

import (
	"context"
	"fmt"
	"login_module/internal/application/dto"
	"login_module/internal/domain/model"
	"login_module/internal/domain/repository"
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

func (s *OAuthService) Login(ctx context.Context, req dto.OAuthDTO) (*dto.OAuthResponse, error) {
	token, err := req.Provider.Exchange(ctx, req.Code)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange code: %s", err.Error())
	}
	res := &dto.OAuthResponse{
		AccessToken:  token.AccessToken,
		TokenType:    token.TokenType,
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry,
	}
	if expiresIn, ok := token.Extra("expires_in").(int); ok {
		res.ExpiresIn = expiresIn
	}
	if refreshTokenExpiresIn, ok := token.Extra("refresh_token_expires_in").(int); ok {
		res.RefreshTokenExpiresIn = refreshTokenExpiresIn
	}
	if idToken, ok := token.Extra("id_token").(string); ok {
		res.IDToken = idToken
	}
	id, err := jwt.ParseIDToken(res.IDToken)
	if err != nil {
		return nil, fmt.Errorf("failed to parse id token: %s", err.Error())
	}
	res.UserUUID = id

	saveToken := model.SaveRefreshToken{
		UserUUID:              res.UserUUID,
		RefreshToken:          res.RefreshToken,
		RefreshTokenExpiresIn: res.RefreshTokenExpiresIn,
	}
	if err := s.redisClient.SaveRefreshToken(ctx, saveToken); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *OAuthService) Logout(ctx context.Context) error {
	userUUID, ok := ctx.Value("userUUID").(string)
	if !ok {
		return fmt.Errorf("userUUID not found in context or not a string")
	}
	err := s.redisClient.DeleteIDToken(ctx, userUUID)
	if err != nil {
		return err
	}
	return nil
}
