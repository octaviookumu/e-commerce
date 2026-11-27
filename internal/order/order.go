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

// entity - private to application
type Order struct {
	ID         string    `json:"id"`
	Status     string    `json:"status"`
	CustomerID string    `json:"customer_id"`
	CreatedAt  time.Time `json:"created_at"`
}



// public
type OrderItemInput struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type PlaceOrderRequest struct {
	CustomerID string           `json:"customer_id"`
	Items      []OrderItemInput `json:"items"` // not showing complexity or internal details to consumer
	// OrderItemInput is an implementation detail we want to keep to ourselves
}
