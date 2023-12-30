package server

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

// Config stores app configuration
type Config struct {
	Address     string `env:"HTTP_SERVER_ADDRESS,default=0.0.0.0:8080"`
	ReadTimeout int    `env:"READ_TIMEOUT,default=5"`
	IdleTimeout int    `env:"IDLE_TIMEOUT,default=30"`
}

// NewConfig reads config from env and creates config struct
func NewConfig() (*Config, error) {
	ctx := context.Background()
	var cfg Config

	if err := envconfig.Process(ctx, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
