package db

import (
	"context"
	"gorm.io/gorm"
	"login_module/internal/domain/model"
	"login_module/internal/domain/repository"
)

type MySQLUserRepository struct {
	db *gorm.DB
}

func NewMySQLUserRepository(db *gorm.DB) repository.UserRepository {
	return &MySQLUserRepository{db: db}
}

func (repo *MySQLUserRepository) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	if err := repo.db.WithContext(ctx).Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *MySQLUserRepository) Save(ctx context.Context, user *model.User) error {
	return repo.db.WithContext(ctx).Create(user).Error
}
