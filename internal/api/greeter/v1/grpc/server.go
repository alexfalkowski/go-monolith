package grpc

import (
	echoer "github.com/alexfalkowski/go-monolith/api/echoer/v1"
	greeter "github.com/alexfalkowski/go-monolith/api/greeter/v1"
	"github.com/alexfalkowski/go-monolith/internal/api/grpcgraph"
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/errors"
	"github.com/alexfalkowski/go-service/v2/net/grpc"
	"github.com/alexfalkowski/go-service/v2/net/grpc/codes"
	"github.com/alexfalkowski/go-service/v2/net/grpc/status"
)

// Register server.
func Register(registrar grpc.ServiceRegistrar, server *Server) {
	greeter.RegisterServiceServer(registrar, server)
}

// NewServer for gRPC.
func NewServer(client echoer.ServiceClient, graph *grpcgraph.Graph) *Server {
	return &Server{client: client, graph: graph}
}

// Server for gRPC.
type Server struct {
	greeter.UnimplementedServiceServer
	client echoer.ServiceClient
	graph  *grpcgraph.Graph
}

// Hello sends a greeting.
func (s *Server) Hello(ctx context.Context, req *greeter.HelloRequest) (*greeter.HelloResponse, error) {
	caller := grpcgraph.ServiceName(greeter.Service_ServiceDesc.ServiceName)
	callee := grpcgraph.ServiceName(echoer.Service_ServiceDesc.ServiceName)

	if !s.graph.CanCall(caller, callee) {
		return nil, status.Errorf(codes.PermissionDenied, "greeter: caller %q is not allowed to call %q", caller, callee)
	}

	ctx = grpcgraph.WithCaller(ctx, caller)

	resp, err := s.client.Echo(ctx, &echoer.EchoRequest{Message: req.GetName()})
	if err != nil {
		return nil, errors.Prefix("greeter", err)
	}

	return &greeter.HelloResponse{Message: "Hello " + resp.GetMessage()}, nil
}
