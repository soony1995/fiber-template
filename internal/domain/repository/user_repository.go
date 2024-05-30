package repository

import (
	"context"
	"login_module/internal/domain/model"
)

type UserRepository interface {
	SaveUser(ctx context.Context, user model.User) error
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
}
