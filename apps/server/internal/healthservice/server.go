package healthservice

import (
	"bxc-fantasy-app/server/gen/protos/go/health/v1/healthv1connect"
	"bxc-fantasy-app/server/internal/healthservice/health"

	"connectrpc.com/connect"
	chi "github.com/go-chi/chi/v5"
)

func RegisterHandlers(router *chi.Mux) {
	healthService := health.NewServer()
	router.Handle("/health/v1/check", connect.NewUnaryHandler(
		healthv1connect.HealthCheckProcedure,
		healthService.Check,
	))
}
