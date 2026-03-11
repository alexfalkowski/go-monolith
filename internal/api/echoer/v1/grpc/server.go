package grpc

import (
	v1 "github.com/alexfalkowski/go-monolith/api/echoer/v1"
	"github.com/alexfalkowski/go-monolith/internal/api/grpcgraph"
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/net/grpc"
	"github.com/alexfalkowski/go-service/v2/net/grpc/codes"
	"github.com/alexfalkowski/go-service/v2/net/grpc/status"
)

// Register server.
func Register(registrar grpc.ServiceRegistrar, server *Server) {
	v1.RegisterServiceServer(registrar, server)
}

// NewServer for gRPC.
func NewServer(graph *grpcgraph.Graph) *Server {
	return &Server{graph: graph}
}

// Server for gRPC.
type Server struct {
	v1.UnimplementedServiceServer
	graph *grpcgraph.Graph
}

// Echo repeats what was sent.
func (s *Server) Echo(ctx context.Context, req *v1.EchoRequest) (*v1.EchoResponse, error) {
	caller, ok := grpcgraph.Caller(ctx)
	if !ok {
		return nil, status.Error(codes.PermissionDenied, "echoer: missing caller service metadata")
	}

	target := grpcgraph.ServiceName(v1.Service_ServiceDesc.ServiceName)
	if !s.graph.CanBeCalledBy(target, caller) {
		return nil, status.Errorf(codes.PermissionDenied, "echoer: caller %q is not allowed to call %q", caller, target)
	}

	return &v1.EchoResponse{Message: req.GetMessage()}, nil
}
