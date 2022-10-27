package models

import "time"

type WebRequestInsertMessage struct {
	Pesan    string `json:"message"`
	Penerima int    `json:"penerima" validate:"required"`
	Pengirim int    `json:"pengirim" validate:"required"`
}

type WebResponInsertMessage struct {
	Pesan      string    `json:"message"`
	Penerima   int       `json:"penerima" validate:"required"`
	Pengirim   int       `json:"pengirim" validate:"required"`
	Status     string    `json:"status"`
	Created_at time.Time `json:"created_at"`
}
