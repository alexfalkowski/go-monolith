package client

import (
	"context"

	"github.com/alexfalkowski/go-service/v2/client"
	"github.com/alexfalkowski/go-service/v2/di"
	"github.com/alexfalkowski/go-service/v2/env"
	"github.com/alexfalkowski/go-service/v2/id"
	"github.com/alexfalkowski/go-service/v2/telemetry/logger"
	"github.com/alexfalkowski/go-service/v2/telemetry/metrics"
	"github.com/alexfalkowski/go-service/v2/telemetry/tracer"
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
	Tracer    *tracer.Tracer
	Meter     *metrics.Meter
	ID        id.Generator
	Client    *client.Config
	Logger    *logger.Logger
	Limiter   *limiter.Client
	UserAgent env.UserAgent
}

// NewClient for grpc.
func NewClient(params Params) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(params.Client.Address,
		grpc.WithClientLogger(params.Logger), grpc.WithClientTracer(params.Tracer),
		grpc.WithClientMetrics(params.Meter), grpc.WithClientRetry(params.Client.Retry),
		grpc.WithClientUserAgent(params.UserAgent), grpc.WithClientTimeout(params.Client.Timeout),
		grpc.WithClientTLS(params.Client.TLS), grpc.WithClientID(params.ID),
		grpc.WithClientLimiter(params.Limiter),
	)

	params.Lifecycle.Append(fx.Hook{
		OnStop: func(_ context.Context) error {
			return conn.Close()
		},
	})

	return conn, err
}
