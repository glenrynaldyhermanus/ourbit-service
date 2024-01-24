// models/order_item_rating.go
package models

import "time"

type OrderItemRating struct {
	ID          int       `json:"id"`
	UUID        string    `json:"uuid"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	IsActive    bool      `json:"is_active"`
	OrderItemID int       `json:"order_item_id"`
	CustomerID  int       `json:"customer_id"`
	Rating      float64   `json:"rating"`
	Comment     string    `json:"comment,omitempty"`
}
