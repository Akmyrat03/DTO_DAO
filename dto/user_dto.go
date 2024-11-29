package dto

type UserRequest struct {
	Username string `json:"username" binding:"required" example:"johndoe"`
	Password string `json:"password" binding:"required" example:"securepassword"`
}

type UserResponse struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}
