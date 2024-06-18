package model

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

func (u *User) Authenticate(password string) bool {
	return u.Password == password
}
