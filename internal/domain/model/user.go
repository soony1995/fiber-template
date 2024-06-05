package model

import (
	"time"
)

type User struct {
	UserUUID     string
	Email        string `json:"email"`
	Password     string
	NickName     string
	Provider     string
	RegisteredAt time.Time
	LastLoginAt  time.Time
}

type SaveRefreshToken struct {
	UserUUID     string `json:"user_uuid"`
	RefreshToken string `json:"refresh_token"`
	Exp          int    `json:"exp"`
}
