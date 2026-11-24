package main

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// payload/request signature
type CreateUserRequest struct {
	Name string `json:"name"`
}

var users = map[string]User{
	"9f21b35f-230d-4fb4-a15d-0bc36ba9a49e": {
		ID:   "9f21b35f-230d-4fb4-a15d-0bc36ba9a49e",
		Name: "Njoki",
	},
	"46c0e632-6e66-4f44-a052-87ff5c2fc609": {
		ID:   "46c0e632-6e66-4f44-a052-87ff5c2fc609",
		Name: "Macharia",
	},
}
