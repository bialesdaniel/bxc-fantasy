package main

import (
	"bxc-fantasy-app/server/internal/healthservice"
	"log"
	"net/http"

	chi "github.com/go-chi/chi/v5"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	log.Println("Starting server...")
	router := chi.NewRouter()

	healthservice.RegisterHandlers(router)

	addr := ":8080"
	log.Printf("Server listening on %s", addr)

	err := http.ListenAndServe(addr, h2c.NewHandler(router, &http2.Server{}))
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
