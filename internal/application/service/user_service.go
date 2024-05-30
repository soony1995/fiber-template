package service

import (
	"context"
	"login_module/internal/application/dto"
	"login_module/internal/domain/model"
	"login_module/internal/domain/repository"

	"github.com/pkg/errors"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetUser(ctx context.Context, email string) (*dto.UserDTO, error) {
	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user")
	}
	userDTO := &dto.UserDTO{
		Email:       user.Email,
		Name:        user.Name,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		NickName:    user.NickName,
		Description: user.Description,
		UserID:      user.UserID,
		AvatarURL:   user.AvatarURL,
		Location:    user.Location,
	}
	return userDTO, nil
}

func (s *UserService) CreateUser(ctx context.Context, userDTO *dto.UserDTO) error {
	user := model.User{
		Email:       userDTO.Email,
		Name:        userDTO.Name,
		FirstName:   userDTO.FirstName,
		LastName:    userDTO.LastName,
		NickName:    userDTO.NickName,
		Description: userDTO.Description,
		UserID:      userDTO.UserID,
		AvatarURL:   userDTO.AvatarURL,
		Location:    userDTO.Location,
	}
	return s.userRepo.SaveUser(ctx, user)
}
