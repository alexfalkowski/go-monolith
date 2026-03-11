package grpcgraph

import "github.com/alexfalkowski/go-service/v2/di"

// Module provides the local in-memory gRPC service graph.
var Module = di.Module(
	di.Constructor(New),
	di.Register(Register),
)
