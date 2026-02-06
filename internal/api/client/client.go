package client

import (
	"context"

	"github.com/alexfalkowski/go-service/v2/config/client"
	"github.com/alexfalkowski/go-service/v2/di"
	"github.com/alexfalkowski/go-service/v2/env"
	"github.com/alexfalkowski/go-service/v2/id"
	"github.com/alexfalkowski/go-service/v2/telemetry/logger"
	"github.com/alexfalkowski/go-service/v2/time"
	"github.com/alexfalkowski/go-service/v2/transport/grpc"
	"github.com/alexfalkowski/go-service/v2/transport/grpc/limiter"
	"go.uber.org/fx"
)

// NewClientLimiter for grpc.
func NewClientLimiter(lc fx.Lifecycle, keys limiter.KeyMap, cfg *client.Config) (*limiter.Client, error) {
	return limiter.NewClientLimiter(lc, keys, cfg.Limiter)
}

// Params for grpc.
type Params struct {
	di.In

	Lifecycle fx.Lifecycle
	ID        id.Generator
	Client    *client.Config
	Logger    *logger.Logger
	Limiter   *limiter.Client
	UserAgent env.UserAgent
}

// NewClient for grpc.
func NewClient(params Params) (*grpc.ClientConn, error) {
	timeout := time.MustParseDuration(params.Client.Timeout)
	keepalivePing := params.Client.Options.Duration("keepalive_ping", timeout)
	keepaliveTimeout := params.Client.Options.Duration("keepalive_timeout", timeout)
	conn, err := grpc.NewClient(params.Client.Address,
		grpc.WithClientLogger(params.Logger), grpc.WithClientRetry(params.Client.Retry),
		grpc.WithClientUserAgent(params.UserAgent), grpc.WithClientID(params.ID),
		grpc.WithClientTimeout(timeout),
		grpc.WithClientKeepalive(keepalivePing, keepaliveTimeout),
		grpc.WithClientTLS(params.Client.TLS),
		grpc.WithClientLimiter(params.Limiter),
	)

	params.Lifecycle.Append(fx.Hook{
		OnStop: func(_ context.Context) error {
			return conn.Close()
		},
	})

	return conn, err
}
