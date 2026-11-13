package v1

import (
	"github.com/alexfalkowski/go-monolith/internal/api/echoer/v1/grpc"
	"github.com/alexfalkowski/go-monolith/internal/api/echoer/v1/http"
	"github.com/alexfalkowski/go-signal/v2"
)

// Module for signal.
var Module = signal.Module(
	grpc.Module,
	http.Module,
)
