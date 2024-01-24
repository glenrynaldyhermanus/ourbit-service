// models/catalogue.go
package models

import "time"

type Catalogue struct {
	ID         int       `json:"id"`
	UUID       string    `json:"uuid"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
	IsActive   bool      `json:"is_active"`
	StoreID    int       `json:"store_id"`
	Name       string    `json:"name,omitempty"`
	PictureURL string    `json:"picture_url,omitempty"`
	ProductID  int       `json:"product_id"`
}
