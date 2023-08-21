package model

type RegisterRequest struct {
	Username string `validate:"required,min=1,max=100"`
	Email    string `validate:"required,email,min=1,max=255"`
	Password string `validate:"required,min=1,max=100"`
}

type LoginRequest struct {
	Email    string `validate:"required,email,min=1,max=255"`
	Password string `validate:"required,min=1,max=100"`
}