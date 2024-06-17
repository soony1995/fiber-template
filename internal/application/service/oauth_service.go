package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"login_module/internal/application/dto"
	"login_module/internal/domain/model"
	"login_module/internal/domain/repository"
	"os"
	"strconv"
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
	token, err := req.Config.Exchange(ctx, req.Code)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange code: %s", err.Error())
	}
	token.Valid()
	tokenExp, err := strconv.Atoi(os.Getenv("HTTP_COOKIE_ACCESS_EXPIRY"))
	if err != nil {
		return nil, fmt.Errorf("failed to exchange code: %s", err.Error())
	}
	refreshTokenExp, err := strconv.Atoi(os.Getenv("HTTP_COOKIE_ACCESS_EXPIRY"))
	if err != nil {
		return nil, fmt.Errorf("failed to exchange code: %s", err.Error())
	}
	// access token + refresh token 을 이용해서 session 을 생성한다.
	// session id를 session
	res := &dto.OAuthResponse{
		UserUUID:              uuid.New().String(),
		AccessToken:           token.AccessToken,
		TokenExpiresIn:        tokenExp,
		TokenType:             token.TokenType,
		RefreshToken:          token.RefreshToken,
		RefreshTokenExpiresIn: refreshTokenExp,
	}
	if idToken, ok := token.Extra("id_token").(string); ok {
		res.IDToken = idToken
	}
	saveToken := model.SaveRefreshToken{
		UserUUID:              res.UserUUID,
		Provider:              req.Provider,
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
