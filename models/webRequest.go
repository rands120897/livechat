package models

type WebRequestData struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"Password" validate:"required"`
}
