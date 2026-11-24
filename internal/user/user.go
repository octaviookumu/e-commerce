package user

import "errors"

// user types, interface, sentinel errors
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// payload/request signature
type CreateUserRequest struct {
	Name string `json:"name"`
}

var (
	ErrNotFound     = errors.New("user not found")
	ErrNameRequired = errors.New("name is required")
)
