package dto

type EmailVerificationDTO struct {
	Email        string `json:"email" form:"email" binding:"required,email"`
	Verification string `json:"verification-code" form:"verification-code" binding:"required"`
}