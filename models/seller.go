// models/seller.go
package models

import "time"

type Seller struct {
	ID            int       `json:"id"`
	UUID          string    `json:"uuid"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
	IsActive      bool      `json:"is_active"`
	FullName      string    `json:"full_name"`
	Email         string    `json:"email"`
	Phone         string    `json:"phone,omitempty"`
	Password      string    `json:"password"`
	IsPhoneActive bool      `json:"is_phone_active"`
	IsEmailActive bool      `json:"is_email_active"`
}
