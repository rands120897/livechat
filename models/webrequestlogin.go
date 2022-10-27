package models

type WebRequestLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"Password" validate:"required"`
}
