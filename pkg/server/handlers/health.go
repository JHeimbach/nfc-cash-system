package handlers

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jheimbach/nfc-cash-system/api"
	"google.golang.org/grpc"
)

type healthServer struct {
}

func RegisterHealthServer(s *grpc.Server) {
	api.RegisterHealthServiceServer(s, &healthServer{})
}

func (h *healthServer) Health(context.Context, *empty.Empty) (*api.HealthStatus, error) {
	return &api.HealthStatus{
		Running: true,
	}, nil
}
