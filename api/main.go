package main

// entrypoint; server setup; create layers; dependency injection

import (
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

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
