package dto

type UserDTO struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	NickName    string `json:"nickname"`
	Description string `json:"description"`
	UserID      string `json:"user_id"`
	AvatarURL   string `json:"avatar_url"`
	Location    string `json:"location"`
}
