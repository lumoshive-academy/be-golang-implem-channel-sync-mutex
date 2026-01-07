package entity

import "time"

type Order struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Total     int64     `json:"total"`
	CreatedAt time.Time `json:"created_at"`
}
