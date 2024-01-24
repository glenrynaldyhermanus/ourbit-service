// models/stock_movement.go
package models

import "time"

type StockMovement struct {
	ID                 int       `json:"id"`
	UUID               string    `json:"uuid"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at,omitempty"`
	IsActive           bool      `json:"is_active"`
	CatalogueVariantID int       `json:"catalogue_variant_id"`
	Movement           int       `json:"movement"`
	OrderNo            string    `json:"order_no,omitempty"`
}
