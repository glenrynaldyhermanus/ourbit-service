// models/business.go
package models

import "time"

type Business struct {
	ID         int       `json:"id"`
	UUID       string    `json:"uuid"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
	IsActive   bool      `json:"is_active"`
	Name       string    `json:"name"`
	Status     string    `json:"status,omitempty"`
	SellerID   int       `json:"seller_id"`
	PictureURL string    `json:"picture_url,omitempty"`
	Username   string    `json:"username"`
	Category   string    `json:"category"`
}
