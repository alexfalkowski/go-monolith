package v1

import (
	"github.com/alexfalkowski/go-monolith/internal/api/echoer/v1/grpc"
	"github.com/alexfalkowski/go-monolith/internal/api/echoer/v1/http"
	"github.com/alexfalkowski/go-service/v2/di"
)

// Module for di.
var Module = di.Module(
	grpc.Module,
	http.Module,
)
