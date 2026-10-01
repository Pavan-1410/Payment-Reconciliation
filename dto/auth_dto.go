package dto

type RegisterRequest struct{
	Name string	`json:"name" binding:"required" example:"Pavan Sonawane"`
	Email string `json:"email" binding:"required,email" example:"pavan@example.com"`
	Password string `json:"password" binding:"required,min=2"  example:"123456"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"pavan@example.com"`
	Password string `json:"password" binding:"required" example:"123456"`
}

type LoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}


type UserResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin"`
}

type ErrorResponse struct {
    Error string `json:"error"`
}