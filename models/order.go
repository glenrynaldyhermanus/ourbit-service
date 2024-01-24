// models/order.go
package models

import "time"

type Order struct {
	ID            int       `json:"id"`
	OrderNo       string    `json:"order_no"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
	IsActive      bool      `json:"is_active"`
	Status        int       `json:"status"`
	CustomerID    int       `json:"customer_id"`
	SellerID      int       `json:"seller_id"`
	ItemCount     int       `json:"item_count"`
	Amount        float64   `json:"amount"`
	ShippingFee   float64   `json:"shipping_fee"`
	AdminFee      float64   `json:"admin_fee"`
	Discount      float64   `json:"discount"`
	Total         float64   `json:"total"`
	PaymentMethod string    `json:"payment_method"`
	PaymentStatus int       `json:"payment_status"`
}
