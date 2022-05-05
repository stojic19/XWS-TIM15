package config

import (
	"os"
)

type Config struct {
	Port           string
	DbHost         string
	DbUsername     string
	DbPassword     string
	DbDatabase     string
	DbNeo4jVersion string
	DbPort         string
}

func NewConfig() *Config {
	return &Config{
		Port:           LookupEnvOrGetDefault("FOLLOWERS_SERVICE_PORT", "8001"),
		DbHost:         LookupEnvOrGetDefault("FOLLOWERS_DB_HOST", "localhost"),
		DbUsername:     LookupEnvOrGetDefault("NEO4J_USER", "neo4j"),
		DbPassword:     LookupEnvOrGetDefault("NEO4J_PASSWORD", "neo4j"),
		DbDatabase:     LookupEnvOrGetDefault("NEO4J_DATABASE", "neo4j"),
		DbPort:         LookupEnvOrGetDefault("FOLLOWERS_DB_PORT", "7687"),
		DbNeo4jVersion: LookupEnvOrGetDefault("NEO4J_VERSION", "4"),
	}
}

func LookupEnvOrGetDefault(key string, defaultValue string) string {
	if env, found := os.LookupEnv(key); !found {
		return defaultValue
	} else {
		return env
	}
}
