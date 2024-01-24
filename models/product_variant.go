// models/product_variant.go
package models

import "time"

type ProductVariant struct {
	ID         int       `json:"id"`
	UUID       string    `json:"uuid"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
	IsActive   bool      `json:"is_active"`
	ProductID  int       `json:"product_id"`
	Name       string    `json:"name"`
	SKU        string    `json:"sku,omitempty"`
	Price      float64   `json:"price"`
	PictureURL string    `json:"picture_url"`
}
