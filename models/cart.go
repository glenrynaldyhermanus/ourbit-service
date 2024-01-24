// models/cart.go
package models

import "time"

type Cart struct {
	ID                 int       `json:"id"`
	UUID               string    `json:"uuid"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at,omitempty"`
	IsActive           bool      `json:"is_active"`
	CustomerID         int       `json:"customer_id"`
	CatalogueVariantID int       `json:"catalogue_variant_id"`
	Quantity           int       `json:"quantity"`
	Notes              string    `json:"notes,omitempty"`
}
