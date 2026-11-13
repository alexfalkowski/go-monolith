package config

import (
	"github.com/alexfalkowski/go-service/v3/config"
	"github.com/alexfalkowski/go-signal/v2"
)

// Module for signal.
var Module = signal.Module(
	signal.Constructor(config.NewConfig[Config]),
	signal.Decorate(decorateConfig),
	signal.Constructor(clientConfig),
	signal.Constructor(healthConfig),
)
