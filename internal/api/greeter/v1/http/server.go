package http

import (
	v1 "github.com/alexfalkowski/go-monolith/api/greeter/v1"
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/errors"
	"github.com/alexfalkowski/go-service/v2/net/http/meta"
	"github.com/alexfalkowski/go-service/v2/net/http/rest"
)

// Register server.
func Register(server *Server) {
	rest.Get("/greeter/v1/hello/{name}", server.Hello)
}

// NewServer for gRPC.
func NewServer(client v1.ServiceClient) *Server {
	return &Server{client: client}
}

// Server for gRPC.
type Server struct {
	client v1.ServiceClient
}

// Hello sends a greeting.
func (s *Server) Hello(ctx context.Context) (*v1.HelloResponse, error) {
	name := meta.Request(ctx).PathValue("name")
	resp, err := s.client.Hello(ctx, &v1.HelloRequest{Name: name})

	return resp, errors.Prefix("greeter", err)
}
