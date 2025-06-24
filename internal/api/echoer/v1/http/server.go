package http

import (
	v1 "github.com/alexfalkowski/go-monolith/api/echoer/v1"
	"github.com/alexfalkowski/go-service/v2/context"
	"github.com/alexfalkowski/go-service/v2/errors"
	"github.com/alexfalkowski/go-service/v2/net/http/meta"
	"github.com/alexfalkowski/go-service/v2/net/http/rest"
)

// Register server.
func Register(server *Server) {
	rest.Get("/echoer/v1/echo/{msg}", server.Echo)
}

// NewServer for http.
func NewServer(client v1.ServiceClient) *Server {
	return &Server{client: client}
}

// Server for http.
type Server struct {
	client v1.ServiceClient
}

// Echo repeats what was sent.
func (s *Server) Echo(ctx context.Context) (*v1.EchoResponse, error) {
	msg := meta.Request(ctx).PathValue("msg")
	resp, err := s.client.Echo(ctx, &v1.EchoRequest{Message: msg})

	return resp, errors.Prefix("echoer", err)
}
