package model

import "time"

// Base struct untuk menyimpan informasi umum
type Base struct {
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}
