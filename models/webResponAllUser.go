package models

import "time"

type WebResponseAllUser struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Created_at time.Time `json:"created_at"`
}
