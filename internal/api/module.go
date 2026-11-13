package api

import (
	echo "github.com/alexfalkowski/go-monolith/internal/api/echoer/v1"
	greeter "github.com/alexfalkowski/go-monolith/internal/api/greeter/v1"
	"github.com/alexfalkowski/go-signal/v2"
)

// Module for signal.
var Module = signal.Module(
	echo.Module,
	greeter.Module,
)
