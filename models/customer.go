package models

import "time"

type Customer struct {
	ID         int       `json:"id"`
	UUID       string    `json:"uuid"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
	IsActive   bool      `json:"is_active"`
	FullName   string    `json:"full_name"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone,omitempty"`
	Password   string    `json:"password"`
	PictureURL string    `json:"picture_url,omitempty"`
	Gender     string    `json:"gender,omitempty"`
	Birthdate  time.Time `json:"birthdate,omitempty"`
}
