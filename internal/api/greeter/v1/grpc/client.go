package grpc

import (
	v1 "github.com/alexfalkowski/go-monolith/api/greeter/v1"
	"github.com/alexfalkowski/go-service/v2/transport/grpc"
)

// NewClient for echo.
func NewClient(conn *grpc.ClientConn) v1.ServiceClient {
	return v1.NewServiceClient(conn)
}
