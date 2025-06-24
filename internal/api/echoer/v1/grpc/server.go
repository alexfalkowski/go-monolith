package grpc

import (
	v1 "github.com/alexfalkowski/go-monolith/api/echoer/v1"
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/net/grpc"
)

// Register server.
func Register(registrar grpc.ServiceRegistrar, server *Server) {
	v1.RegisterServiceServer(registrar, server)
}

// NewServer for gRPC.
func NewServer() *Server {
	return &Server{}
}

// Server for gRPC.
type Server struct {
	v1.UnimplementedServiceServer
}

// Echo repeats what was sent.
func (s *Server) Echo(_ context.Context, req *v1.EchoRequest) (*v1.EchoResponse, error) {
	return &v1.EchoResponse{Message: req.GetMessage()}, nil
}
