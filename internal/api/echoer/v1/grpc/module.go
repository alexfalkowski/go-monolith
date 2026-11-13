package grpc

import "github.com/alexfalkowski/go-signal/v2"

// Module for signal.
var Module = signal.Module(
	signal.Constructor(NewClient),
	signal.Constructor(NewServer),
	signal.Register(Register),
)
