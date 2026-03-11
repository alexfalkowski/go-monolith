package api

import (
	"github.com/alexfalkowski/go-monolith/internal/api/client"
	echo "github.com/alexfalkowski/go-monolith/internal/api/echoer/v1"
	greeter "github.com/alexfalkowski/go-monolith/internal/api/greeter/v1"
	"github.com/alexfalkowski/go-monolith/internal/api/grpcgraph"
	"github.com/alexfalkowski/go-service/v2/di"
)

// Module for di.
var Module = di.Module(
	client.Module,
	echo.Module,
	greeter.Module,
	grpcgraph.Module,
)
