package repository

import (
	"context"
	"login_module/internal/domain/model"
)

type UserRepository interface {
	FindByUsername(ctx context.Context, username string) (*model.User, error)
	Save(ctx context.Context, user *model.User) error
}
