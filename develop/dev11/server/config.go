package server

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

// Stores app configuration
type Config struct {
	Address     string `env:"ADDRESS" envDefault:"8080"`
	ReadTimeout int    `env:"READ_TIMEOUT" envDefault:"3"`
	IdleTimeout int    `env:"IDLE_TIMEOUT" envDefault:"10"`
}

// Parses a .env file and return a Config
func NewConfig() (*Config, error) {
	var cfg Config

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("%+v", err)
	}

	return &cfg, nil
}
