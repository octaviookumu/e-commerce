package user

// business logic. storage functions

import (
	"context" // will help with cancelling requests, adding timeouts
	"fmt"

	"github.com/google/uuid"
)

// Repository: the contract - similar to a repository interface
type Repository interface {
	Create(ctx context.Context, u User) (User, error)  // create a user
	ByID(ctx context.Context, id string) (User, error) // get a user by id
	List(ctx context.Context) ([]User, error)          // list users
} // Any Go type with these three methods automatically satisfies this interface

// Service: business logic
type Service struct {
	repo Repository // dependencies. will dependency inject
}

// dependency injection - this is a constructor
// NewService(repo) builds a Service and saves the passed repository in its repo field.
// will be used as:
// repo := user.NewMemoryRepository()
// service := user.NewService(repo)
func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// methods for the service
func (s *Service) Get(ctx context.Context, id string) (User, error) {
	return s.repo.ByID(ctx, id)
}

func (s *Service) List(ctx context.Context) ([]User, error) {
	return s.repo.List(ctx)
}

func (s *Service) Create(ctx context.Context, payload CreateUserRequest) (User, error) {
	if payload.Name == "" {
		return User{}, ErrNameRequired
	}

	newID := uuid.New().String()
	newUser := User{
		ID:   newID,
		Name: payload.Name,
	}

	if _, err := s.repo.Create(ctx, newUser); err != nil {
		return User{}, fmt.Errorf("create user: %w", err)
		// using a custom error because I don't want to send this error to the user
	}

	return newUser, nil

}
