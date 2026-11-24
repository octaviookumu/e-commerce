package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type server struct {
	addr string
}

func (s *server) getUsersListHandler(w http.ResponseWriter, r *http.Request) {
	usersList := []User{}

	for _, user := range users {
		usersList = append(usersList, user)
	}
	writeJSON(w, http.StatusOK, usersList)
}

func (s *server) createUserHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	decodeError := json.NewDecoder(r.Body).Decode(&req) // JSON -> Go (Decode)
	if decodeError != nil {
		handleError(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Name == "" {
		handleError(w, "name is required", http.StatusBadRequest)
		return
	}

	newID := uuid.New().String()
	newUser := User{
		ID:   newID,
		Name: req.Name,
	}
	users[newID] = newUser

	writeJSON(w, http.StatusCreated, newUser)
}

func (s *server) getUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	// validation
	user, exists := users[id]
	if !exists {
		handleError(w, "user not found", http.StatusNotFound)
	}

	writeJSON(w, http.StatusOK, user)
}

func main() {
	// The & means “take the address of this value.”
	// So srv is a pointer to a server, written as *server
	srv := &server{addr: ":8080"}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /users", srv.getUsersListHandler)
	mux.HandleFunc("POST /users", srv.createUserHandler)
	mux.HandleFunc("GET /users/{id}", srv.getUserHandler)

	log.Println("Server is running...")

	http.ListenAndServe(srv.addr, mux)
}
