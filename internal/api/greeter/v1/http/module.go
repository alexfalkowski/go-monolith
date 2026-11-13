package http

import "github.com/alexfalkowski/go-signal/v2"

// Module for signal.
var Module = signal.Module(
	signal.Constructor(NewServer),
	signal.Register(Register),
)
