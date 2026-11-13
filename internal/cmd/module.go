package cmd

import (
	"github.com/alexfalkowski/go-monolith/internal/api"
	"github.com/alexfalkowski/go-monolith/internal/api/client"
	"github.com/alexfalkowski/go-monolith/internal/config"
	"github.com/alexfalkowski/go-monolith/internal/health"
	"github.com/alexfalkowski/go-service/v3/module"
	"github.com/alexfalkowski/go-signal/v2"
)

// Module for signal.
var Module = signal.Module(
	module.Server,
	config.Module,
	health.Module,
	client.Module,
	api.Module,
)
