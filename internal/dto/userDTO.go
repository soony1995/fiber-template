package dto

// UserDTO 사용자 정보 전달을 위한 DTO 정의
type UserDTO struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
