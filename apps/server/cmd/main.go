package main

import (
	"log"
	"net/http"
	"os"

	"bxc-fantasy-app/server/internal/healthservice"

	chi "github.com/go-chi/chi/v5"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port

	log.Println("Starting server...")
	router := chi.NewRouter()

	healthservice.RegisterHandlers(router)

	log.Printf("Server listening on %s", addr)

	err := http.ListenAndServe(addr, h2c.NewHandler(router, &http2.Server{}))
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
