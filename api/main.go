package main

// entrypoint; server setup; create layers; dependency injection

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/octaviookumu/myproject/internal/order"
	"github.com/octaviookumu/myproject/internal/product"
	"github.com/octaviookumu/myproject/internal/user"
)

func main() {
	ctx := context.Background()

	// set up layers: STORAGE -> SERVICE -> TRANSPORT
	userRepo := user.NewMemoryRepository()
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	orderRepo := order.NewMemoryRepository()
	orderService := order.NewService(orderRepo)
	orderHandler := order.NewHandler(orderService)

	productRepo := product.NewMemoryRepository()
	productService := product.NewService(productRepo)
	productHandler := product.NewHandler(productService)

	mux := http.NewServeMux()
	userHandler.Register(mux) // register our handlers into global mux
	orderHandler.Register(mux)
	productHandler.Register(mux)

	if err := product.Seed(ctx, productRepo); err != nil {
		log.Fatal(err)
	}

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
