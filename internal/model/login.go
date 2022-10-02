package model

type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type PhoneLoginRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

type PhoneLoginCode struct {
	VerificationCode string `json:"verification_code" binding:"required"`
}
