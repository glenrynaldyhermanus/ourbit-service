// models/wishlist.go
package models

import "time"

type Wishlist struct {
	ID          int       `json:"id"`
	UUID        string    `json:"uuid"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	IsActive    bool      `json:"is_active"`
	CustomerID  int       `json:"customer_id"`
	CatalogueID int       `json:"catalogue_id"`
}
