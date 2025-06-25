package config

import (
	"github.com/alexfalkowski/go-monolith/internal/health"
	"github.com/alexfalkowski/go-service/v2/client"
	"github.com/alexfalkowski/go-service/v2/config"
)

// Config for the service.
type Config struct {
	Client         *client.Config `yaml:"client,omitempty" json:"client,omitempty" toml:"client,omitempty"`
	Health         *health.Config `yaml:"health,omitempty" json:"health,omitempty" toml:"health,omitempty"`
	*config.Config `yaml:",inline" json:",inline" toml:",inline"`
}

func decorateConfig(cfg *Config) *config.Config {
	return cfg.Config
}

func clientConfig(cfg *Config) *client.Config {
	return cfg.Client
}

func healthConfig(cfg *Config) *health.Config {
	return cfg.Health
}
