package config

import (
	"os"
)

type Config struct {
	Port                string
	DbHost              string
	DbUsername          string
	DbPassword          string
	DbDatabase          string
	DbNeo4jVersion      string
	DbPort              string
	UsersHost           string
	UsersPort           string
	NatsHost            string
	NatsPort            string
	NatsUser            string
	NatsPass            string
	BlockCommandSubject string
	BlockReplySubject   string
}

func NewConfig() *Config {
	return &Config{
		Port:                LookupEnvOrGetDefault("FOLLOWERS_SERVICE_PORT", "8001"),
		DbHost:              LookupEnvOrGetDefault("FOLLOWERS_DB_HOST", "localhost"),
		DbUsername:          LookupEnvOrGetDefault("NEO4J_USER", "neo4j"),
		DbPassword:          LookupEnvOrGetDefault("NEO4J_PASSWORD", "neo4j"),
		DbDatabase:          LookupEnvOrGetDefault("NEO4J_DATABASE", "neo4j"),
		DbPort:              LookupEnvOrGetDefault("FOLLOWERS_DB_PORT", "7687"),
		DbNeo4jVersion:      LookupEnvOrGetDefault("NEO4J_VERSION", "4"),
		UsersHost:           LookupEnvOrGetDefault("USERS_SERVICE_HOST", "localhost"),
		UsersPort:           LookupEnvOrGetDefault("USERS_PORT", "9090"),
		NatsHost:            LookupEnvOrGetDefault("NATS_HOST", "localhost"),
		NatsPort:            LookupEnvOrGetDefault("NATS_PORT", "4222"),
		NatsUser:            LookupEnvOrGetDefault("NATS_USER", "ruser"),
		NatsPass:            LookupEnvOrGetDefault("NATS_PASS", "T0pS3cr3t"),
		BlockCommandSubject: LookupEnvOrGetDefault("BLOCK_COMMAND_SUBJECT", "block.command"),
		BlockReplySubject:   LookupEnvOrGetDefault("BLOCK_REPLY_SUBJECT", "block.reply"),
	}
}

func LookupEnvOrGetDefault(key string, defaultValue string) string {
	if env, found := os.LookupEnv(key); !found {
		return defaultValue
	} else {
		return env
	}
}
