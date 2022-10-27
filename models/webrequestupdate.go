package models

type WebRequestUpdate struct {
	Id       int    `json:"id"`
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}
