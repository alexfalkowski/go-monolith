package grpc

import (
	echoer "github.com/alexfalkowski/go-monolith/api/echoer/v1"
	greeter "github.com/alexfalkowski/go-monolith/api/greeter/v1"
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/errors"
	"github.com/alexfalkowski/go-service/v2/net/grpc"
)

// Register server.
func Register(registrar grpc.ServiceRegistrar, server *Server) {
	greeter.RegisterServiceServer(registrar, server)
}

// NewServer for gRPC.
func NewServer(client echoer.ServiceClient) *Server {
	return &Server{client: client}
}

// Server for gRPC.
type Server struct {
	greeter.UnimplementedServiceServer
	client echoer.ServiceClient
}

// Hello sends a greeting.
func (s *Server) Hello(ctx context.Context, req *greeter.HelloRequest) (*greeter.HelloResponse, error) {
	resp, err := s.client.Echo(ctx, &echoer.EchoRequest{Message: req.GetName()})
	if err != nil {
		return nil, errors.Prefix("greeter", err)
	}

	return &greeter.HelloResponse{Message: "Hello " + resp.GetMessage()}, nil
}
