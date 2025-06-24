package cmd

import (
	"github.com/alexfalkowski/go-monolith/internal/api"
	"github.com/alexfalkowski/go-monolith/internal/api/client"
	"github.com/alexfalkowski/go-monolith/internal/config"
	"github.com/alexfalkowski/go-monolith/internal/health"
	"github.com/alexfalkowski/go-service/v2/di"
	"github.com/alexfalkowski/go-service/v2/module"
)

// Module for di.
var Module = di.Module(
	module.Server,
	config.Module,
	health.Module,
	client.Module,
	api.Module,
)
