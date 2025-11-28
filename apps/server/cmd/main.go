package main

import (
	"log"
	"net/http"
	"net/url"
	"os"

	"bxc-fantasy-app/server/internal/healthservice"

	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port

	router := http.NewServeMux()
	healthservice.RegisterHandlers(router)

	c := cors.New(cors.Options{
		AllowOriginFunc: func(origin string) bool {
			u, err := url.Parse(origin)
			if err != nil {
				return false
			}
			// TODO: Productionalize this.
			return u.Hostname() == "localhost" || u.Hostname() == "127.0.0.1"
		},
		AllowedMethods:   []string{"POST"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(h2c.NewHandler(router, &http2.Server{}))
	log.Printf("Server listening on %s", addr)

	err := http.ListenAndServe(addr, handler)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
