package order

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Repository interface {
	Create(context.Context, Order) error
	ByID(context.Context, string) (Order, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// service methods
func (s *Service) PlaceOrder(ctx context.Context, req PlaceOrderRequest) (Order, error) {
	if req.CustomerID == "" {
		return Order{}, ErrCustomerRequired
	}

	newID := uuid.New().String()
	order := Order{
		ID:         newID,
		CustomerID: req.CustomerID,
		Status:     string(StatusPending),
		CreatedAt:  time.Now().UTC(),
	}

	if err := s.repo.Create(ctx, order); err != nil {
		return Order{}, err
	}

	return order, nil
}

func (s *Service) GetOrder(ctx context.Context, id string) (Order, error) {
	order, err := s.repo.ByID(ctx, id)
	if err != nil {
		return Order{}, err
	}

	return order, nil
}
