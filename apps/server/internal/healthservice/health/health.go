package health

import (
	"bxc-fantasy-app/server/gen/protos/go/health/v1/healthv1connect"
	healthv1 "bxc-fantasy-app/server/gen/protos/go/health/v1"
	"context"
	"log"
	"time"

	"connectrpc.com/connect"
)

type Server struct {
	healthv1connect.UnimplementedHealthHandler
	startTime time.Time
}

func NewServer() *Server {
	return &Server{
		UnimplementedHealthHandler: healthv1connect.UnimplementedHealthHandler{},
		startTime: time.Now(),
	}
}

func (s *Server) Check(
	ctx context.Context,
	req *connect.Request[healthv1.CheckRequest],
) (*connect.Response[healthv1.CheckResponse], error) {
	log.Println("Received health check request.")
	res := connect.NewResponse(&healthv1.CheckResponse{
		UptimeSeconds: int32(time.Since(s.startTime).Seconds()),
	})
	return res, nil
}
