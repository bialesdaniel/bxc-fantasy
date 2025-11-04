package healthservice

import (
	"context"
	"log"
	"time"

	healthv1 "bxc-fantasy-app/server/gen/protos/go/health/v1"

	"connectrpc.com/connect"
)

func (s *Server) Check(ctx context.Context, req *connect.Request[healthv1.CheckRequest]) (*connect.Response[healthv1.CheckResponse], error) {
	log.Println("Received health check request.")
	return connect.NewResponse(&healthv1.CheckResponse{
		UptimeSeconds: int32(time.Since(s.startTime).Seconds()),
	}), nil
}
