package dto

// All the DTOS that are recieved from HTTP requests

type InLogin struct {
	Email string `json:"email" binding:"required,email"`
	Senha string `json:"senha" binding:"required"`
}
