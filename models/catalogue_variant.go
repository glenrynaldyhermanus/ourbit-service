// models/catalogue_variant.go
package models

import "time"

type CatalogueVariant struct {
	ID               int       `json:"id"`
	UUID             string    `json:"uuid"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`
	IsActive         bool      `json:"is_active"`
	CatalogueID      int       `json:"catalogue_id"`
	Name             string    `json:"name,omitempty"`
	SKU              string    `json:"sku,omitempty"`
	Price            float64   `json:"price,omitempty"`
	PictureURL       string    `json:"picture_url,omitempty"`
	Stock            int       `json:"stock,omitempty"`
	ProductVariantID int       `json:"product_variant_id,omitempty"`
}
