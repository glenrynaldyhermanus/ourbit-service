// models/store.go
package models

import (
	"ourbit-service/models/utils"
	"time"
)

type Store struct {
	ID         int              `json:"id"`
	UUID       string           `json:"uuid"`
	CreatedAt  time.Time        `json:"created_at"`
	UpdatedAt  time.Time        `json:"updated_at,omitempty"`
	IsActive   bool             `json:"is_active"`
	Name       string           `json:"name"`
	BusinessID int              `json:"business_id"`
	SellerID   int              `json:"seller_id"`
	Address    string           `json:"address,omitempty"`
	Coordinate utils.Coordinate `json:"coordinate,omitempty"`
}
