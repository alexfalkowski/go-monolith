package health

import (
	"github.com/alexfalkowski/go-health/v2/checker"
	"github.com/alexfalkowski/go-health/v2/server"
	echoer "github.com/alexfalkowski/go-monolith/api/echoer/v1"
	greeter "github.com/alexfalkowski/go-monolith/api/greeter/v1"
	"github.com/alexfalkowski/go-service/v2/env"
	"github.com/alexfalkowski/go-service/v2/errors"
	"github.com/alexfalkowski/go-service/v2/health"
	"github.com/alexfalkowski/go-service/v2/time"
)

func register(name env.Name, srv *server.Server, cfg *Config) {
	d := time.MustParseDuration(cfg.Duration)
	regs := health.Registrations{
		server.NewRegistration("noop", d, checker.NewNoopChecker()),
		server.NewOnlineRegistration(d, d),
	}

	srv.Register(name.String(), regs...)
	srv.Register(echoer.Service_ServiceDesc.ServiceName, regs[0])
	srv.Register(greeter.Service_ServiceDesc.ServiceName, regs[0])
}

func httpHealthObserver(name env.Name, server *server.Server) error {
	return server.Observe(name.String(), "healthz", "online")
}

func httpLivenessObserver(name env.Name, server *server.Server) error {
	return server.Observe(name.String(), "livez", "noop")
}

func httpReadinessObserver(name env.Name, server *server.Server) error {
	return server.Observe(name.String(), "readyz", "noop")
}

func grpcObserver(server *server.Server) error {
	return errors.Join(
		server.Observe(echoer.Service_ServiceDesc.ServiceName, "grpc", "noop"),
		server.Observe(greeter.Service_ServiceDesc.ServiceName, "grpc", "noop"),
	)
}
