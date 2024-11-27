package dto

type UserRequest struct {
	Username string `json:"username" validate:"required, min=4, max=20"`
	Password string `json:"password" validate:"required, min=6"`
}

type UserResponse struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}
