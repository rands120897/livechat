package models

import "time"

type Message struct {
	Id         int
	Pesan      string
	Penerima   int
	Pengirim   int
	Status     string
	Created_at time.Time
}
