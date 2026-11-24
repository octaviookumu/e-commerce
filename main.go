package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

type server struct {
	addr string
}

// Here, *server means the method belongs to a pointer to a server.
// The variable s is that pointer. Go uses this method to let server handle HTTP requests.
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// GET
	if r.Method == "GET" && r.URL.Path == "/users" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		usersList := []User{}

		for _, user := range users {
			usersList = append(usersList, user)
		}
		json.NewEncoder(w).Encode(usersList) // Go -> JSON (Encode)
		return
	}

	// POST /users
	if r.Method == "POST" && r.URL.Path == "/users" {
		// payload/request signature
		var req CreateUserRequest
		decodeError := json.NewDecoder(r.Body).Decode(&req) // JSON -> Go (Decode)
		if decodeError != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}

		if req.Name == "" {
			http.Error(w, "name is required", http.StatusBadRequest)
			return
		}

		newID := uuid.New().String()
		newUser := User{
			ID:   newID,
			Name: req.Name,
		}
		users[newID] = newUser

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newUser)
		return
	}

	// GET /users/{id}
	if r.Method == "GET" {
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) == 3 && parts[1] == "users" {
			id := parts[2]

			user, exists := users[id]
			if !exists {
				http.Error(w, "user not found", http.StatusNotFound)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
		}
		return
	}

	http.Error(w, "not found", http.StatusNotFound)

}

func main() {
	// The & means “take the address of this value.”
	// So srv is a pointer to a server, written as *server
	srv := &server{addr: ":8080"}

	log.Println("Server is running...")

	http.ListenAndServe(srv.addr, srv)
}
