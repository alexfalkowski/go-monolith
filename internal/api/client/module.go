package client

import "github.com/alexfalkowski/go-service/v2/di"

// Module for di.
var Module = di.Module(
	di.Constructor(NewClient),
)
