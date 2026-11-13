package health

import "github.com/alexfalkowski/go-signal/v2"

// Module for signal.
var Module = signal.Module(
	signal.Register(register),
)
