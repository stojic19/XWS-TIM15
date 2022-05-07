package config

import (
	"os"
)

type Config struct {
	Port string
}

func NewConfig() *Config {
	return &Config{
		Port: LookupEnvOrGetDefault("POSTS_SERVICE_PORT", "8002"),
	}
}

func LookupEnvOrGetDefault(key string, defaultValue string) string {
	if env, found := os.LookupEnv(key); !found {
		return defaultValue
	} else {
		return env
	}
}
