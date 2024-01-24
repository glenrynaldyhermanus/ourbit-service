// models/product.go
package models

import "time"

type Product struct {
	ID         int       `json:"id"`
	UUID       string    `json:"uuid"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
	IsActive   bool      `json:"is_active"`
	Name       string    `json:"name"`
	BusinessID int       `json:"business_id"`
	SellerID   int       `json:"seller_id"`
	PictureURL string    `json:"picture_url"`
}
