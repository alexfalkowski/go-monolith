package client

import (
	"github.com/alexfalkowski/go-service/v3/config/client"
	"github.com/alexfalkowski/go-service/v3/context"
	"github.com/alexfalkowski/go-service/v3/env"
	"github.com/alexfalkowski/go-service/v3/id"
	"github.com/alexfalkowski/go-service/v3/telemetry/logger"
	"github.com/alexfalkowski/go-service/v3/transport/grpc"
	"github.com/alexfalkowski/go-service/v3/transport/grpc/breaker"
	"github.com/alexfalkowski/go-service/v3/transport/grpc/limiter"
	"github.com/alexfalkowski/go-service/v3/transport/grpc/retry"
	"github.com/alexfalkowski/go-signal/v2"
)

// NewClientLimiter for grpc.
func NewClientLimiter(lc signal.Lifecycle, keys limiter.KeyMap, cfg *client.Config) (*limiter.Client, error) {
	return limiter.NewClient(lc, keys, cfg.Limiter)
}

// Params for grpc.
type Params struct {
	signal.In

	Lifecycle signal.Lifecycle
	ID        id.Generator
	Client    *client.Config
	Logger    *logger.Logger
	Limiter   *limiter.Client
	UserAgent env.UserAgent
}

// NewClient for grpc.
func NewClient(params Params) (*grpc.ClientConn, error) {
	keepalivePing := params.Client.Options.Duration("keepalive_ping", params.Client.Timeout)
	keepaliveTimeout := params.Client.Options.Duration("keepalive_timeout", params.Client.Timeout)
	conn, err := grpc.NewClient(params.Client.Address,
		grpc.WithClientLogger(params.Logger),
		grpc.WithClientBreaker(breaker.NewConfig(params.Client.Breaker)),
		grpc.WithClientRetry(retry.NewConfig(params.Client.Retry)),
		grpc.WithClientUserAgent(params.UserAgent), grpc.WithClientID(params.ID),
		grpc.WithClientTimeout(params.Client.Timeout),
		grpc.WithClientKeepalive(keepalivePing, keepaliveTimeout),
		grpc.WithClientTLS(params.Client.TLS),
		grpc.WithClientLimiter(params.Limiter),
	)

	params.Lifecycle.Append(signal.Hook{
		OnStop: func(_ context.Context) error {
			return conn.Close()
		},
	})

	return conn, err
}
