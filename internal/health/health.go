package health

import (
	"github.com/alexfalkowski/go-health/v2/checker"
	"github.com/alexfalkowski/go-health/v2/server"
	echoer "github.com/alexfalkowski/go-monolith/api/echoer/v1"
	greeter "github.com/alexfalkowski/go-monolith/api/greeter/v1"
	"github.com/alexfalkowski/go-service/v2/env"
	"github.com/alexfalkowski/go-service/v2/errors"
	"github.com/alexfalkowski/go-service/v2/health"
)

func register(name env.Name, srv *server.Server, cfg *Config) error {
	regs := health.Registrations{
		server.NewRegistration("noop", cfg.Duration.Duration(), checker.NewNoopChecker()),
		server.NewOnlineRegistration(cfg.Timeout.Duration(), cfg.Duration.Duration()),
	}
	n := name.String()

	srv.Register(n, regs...)
	srv.Register(echoer.Service_ServiceDesc.ServiceName, regs[0])
	srv.Register(greeter.Service_ServiceDesc.ServiceName, regs[0])

	return errors.Join(
		srv.Observe(n, "healthz", "online"),
		srv.Observe(n, "livez", "noop"),
		srv.Observe(n, "readyz", "noop"),
		srv.Observe(echoer.Service_ServiceDesc.ServiceName, "grpc", "noop"),
		srv.Observe(greeter.Service_ServiceDesc.ServiceName, "grpc", "noop"),
	)
}
