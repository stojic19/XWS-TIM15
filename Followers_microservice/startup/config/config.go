package config

import (
	"os"
)

type Config struct {
	Port string
}

func NewConfig() *Config {
	return &Config{
		Port: os.Getenv("FOLLOWERS_SERVICE_PORT"),
	}
}
