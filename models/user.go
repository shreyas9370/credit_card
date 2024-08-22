package models

import "time"

type User struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"-"`
	CardNumber string    `json:"card_number"`
	CreatedAt  time.Time `json:"created_at"`
}
