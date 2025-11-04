package healthservice

import (
	"net/http"
	"time"

	"bxc-fantasy-app/server/gen/protos/go/health/v1/healthv1connect"
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

func RegisterHandlers(router *http.ServeMux) {
	healthService := New()
	path, handler := healthv1connect.NewHealthHandler(healthService)
	router.Handle(path, handler)
}
