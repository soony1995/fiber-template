package dto

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Session  string `json:"session"`
}
