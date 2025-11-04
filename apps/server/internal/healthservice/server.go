package healthservice

import (
	"time"

	"bxc-fantasy-app/server/gen/protos/go/health/v1/healthv1connect"

	"connectrpc.com/connect"
	chi "github.com/go-chi/chi/v5"
)

// Server implements the health.v1.Health service.
type Server struct {
	healthv1connect.UnimplementedHealthHandler
	startTime time.Time
}

// New returns a new health service server.
func New() *Server {
	return &Server{
		startTime: time.Now(),
	}
}

func RegisterHandlers(router *chi.Mux) {
	healthService := New()
	router.Handle("/health/v1/check", connect.NewUnaryHandler(
		healthv1connect.HealthCheckProcedure,
		healthService.Check,
	))
}
