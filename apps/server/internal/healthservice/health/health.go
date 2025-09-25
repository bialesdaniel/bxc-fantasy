package health

import (
	healthv1 "bxc-fantasy-app/server/gen/protos/go/health/v1"
	"context"
	"log"
	"time"

	"connectrpc.com/connect"
)

type Server struct {
	startTime time.Time
}

func NewService() *Server {
	return &Server{
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
