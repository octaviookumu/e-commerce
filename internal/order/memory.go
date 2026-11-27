package order

import (
	"context"
	"sync"
)

type MemoryRepository struct{
	mu sync.RWMutex
}

func NewMemoryRepository()*MemoryRepository{
	return &MemoryRepository{}
}

func (r *MemoryRepository) Create(_ context.Context, order Order)error{

}

func (r *MemoryRepository) ByID(_ context.Context, id string)(Order, error){

}

