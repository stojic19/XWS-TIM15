package config

import (
	"os"
)

type Config struct {
	Port            string
	JobOffersDbPort string
	JobOffersDbHost string
}

func NewConfig() *Config {
	return &Config{
		Port:            LookupEnvOrGetDefault("JOB_OFFERS_SERVICE_PORT", "8003"),
		JobOffersDbHost: LookupEnvOrGetDefault("JOB_OFFERS_DB_HOST", "localhost"),
		JobOffersDbPort: LookupEnvOrGetDefault("JOB_OFFERS_DB_PORT", "27017"),
	}
}

func LookupEnvOrGetDefault(key string, defaultValue string) string {
	if env, found := os.LookupEnv(key); !found {
		return defaultValue
	} else {
		return env
	}
}
