package dtos

type CreateUserDtoResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	UserName string `json:"username"`
}
