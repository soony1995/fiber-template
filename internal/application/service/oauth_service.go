package service

import (
	"context"
	"fmt"
	"login_module/internal/application/dto"
	"login_module/internal/domain/model"
	"login_module/internal/domain/repository"
	"login_module/pkg/util"
)

type OAuthService struct {
	redisClient repository.AuthRepository
}

func NewOAuthService(r repository.AuthRepository) *OAuthService {
	return &OAuthService{
		redisClient: r,
	}
}

func (s *OAuthService) Login(c context.Context, m dto.OAuthDTO) (err error) {
	// check m.UserUuid in mysql server
	// if not exist redirect register page
	exp, err := util.GetenvInt("HTTP_COOKIE_REFRESH_EXPIRY")
	if err != nil {
		return err
	}
	SaveRefreshToken := model.SaveRefreshToken{
		UserUUID:     m.UserUUID,
		RefreshToken: m.Token.RefreshToken,
		Exp:          exp,
	}
	if err := s.redisClient.SaveRefreshToken(c, SaveRefreshToken); err != nil {
		return err
	}
	return
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
