package main

// entrypoint; server setup; create layers; dependency injection

import (
	"log"
	"net/http"
	"time"

	"github.com/octaviookumu/myproject/internal/user"
)

func main() {
	// set up layers: STORAGE -> SERVICE -> TRANSPORT
	userRepo := user.NewMemoryRepository()
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	mux := http.NewServeMux()
	userHandler.Register(mux) // register our handlers into global mux

	srv := &http.Server{
		Addr:        ":8080",
		Handler:     mux,
		ReadTimeout: 5 * time.Second,
	}

	log.Println("Server is running...")

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("Server failed to start", "error", err)
	}
}
