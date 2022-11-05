package dtos

type LoginResponse struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}
