package product

import (
	"errors"
	"time"
)

// entity
type Product struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"` // in cents, the lowest possible unit
	Stock     int       `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
}

var (
	ErrNotFound = errors.New("product not found")
)
