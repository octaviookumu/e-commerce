package order

import (
	"errors"
	"time"
)

var (
	ErrCustomerRequired = errors.New("customer is required")
	ErrNotFound         = errors.New("order not found")
)

type Status string

const StatusPending Status = "pending"

type Order struct {
	ID         string    `json:"id"`
	Status     string    `json:"status"`
	CustomerID string    `json:"customer_id"`
	CreatedAt  time.Time `json:"created_at"`
}

type PlaceOrderRequest struct {
	CustomerID string `json:"customer_id"`
}
