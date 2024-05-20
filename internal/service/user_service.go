package service

import "fmt"

// UserService 인터페이스 정의
type UserService interface {
    GetUser(id int) string
}

type userService struct{}

func NewUserService() UserService {
    return &userService{}
}

func (u *userService) GetUser(id int) string {
    return fmt.Sprintf("User ID: %d", id)
}
