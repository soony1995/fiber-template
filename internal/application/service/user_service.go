package service

import (
	"login_module/internal/domain/repository"
)

type UserService struct {
	userRepo repository.AuthRepository
}

func NewUserService(userRepo repository.AuthRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) Login() error {
	return nil
}
