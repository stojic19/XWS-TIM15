package config

import (
	"os"
)

type Config struct {
	Port        string
	PostsDbPort string
	PostsDbHost string
}

func NewConfig() *Config {
	return &Config{
		Port:        LookupEnvOrGetDefault("POSTS_SERVICE_PORT", "8002"),
		PostsDbHost: LookupEnvOrGetDefault("POSTS_DB_HOST", "localhost"),
		PostsDbPort: LookupEnvOrGetDefault("POSTS_DB_PORT", "27017"),
	}
}

func LookupEnvOrGetDefault(key string, defaultValue string) string {
	if env, found := os.LookupEnv(key); !found {
		return defaultValue
	} else {
		return env
	}
}
