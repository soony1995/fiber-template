package repository

import (
	"context"
	"login_module/internal/domain/model"
)

type AuthRepository interface {
	SaveRefreshToken(c context.Context, token model.SaveRefreshToken) error
	DeleteIDToken(c context.Context, userUUID string) error
	GetUserByUserUUID(c context.Context, email string) (model.User, error)
}
