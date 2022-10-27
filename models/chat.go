package models

import "time"

type Chat struct {
	Id          int       `json:"id"`
	Pesan       string    `json:"pesan"`
	Id_pengirim int       `json:"id_pengirim"`
	Id_penerima int       `json:"id_penerima"`
	Status      bool      `json:"status"`
	Created_at  time.Time `json:"created_at"`
}
