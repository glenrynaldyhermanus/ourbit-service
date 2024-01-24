// models/product_picture.go
package models

import "time"

type ProductPicture struct {
	ID        int       `json:"id"`
	UUID      string    `json:"uuid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	IsActive  bool      `json:"is_active"`
	URL       string    `json:"url"`
	IsMain    bool      `json:"is_main"`
	ProductID int       `json:"product_id"`
}
