package config

import (
	"os"
)

type Config struct {
	Port          string
	PostsDbPort   string
	PostsDbHost   string
	FollowersHost string
	FollowersPort string
	UsersHost     string
	UsersPort     string
}

func NewConfig() *Config {
	return &Config{
		Port:          LookupEnvOrGetDefault("POSTS_SERVICE_PORT", "8002"),
		PostsDbHost:   LookupEnvOrGetDefault("POSTS_DB_HOST", "localhost"),
		PostsDbPort:   LookupEnvOrGetDefault("POSTS_DB_PORT", "27017"),
		FollowersHost: LookupEnvOrGetDefault("FOLLOWERS_SERVICE_HOST", "localhost"),
		FollowersPort: LookupEnvOrGetDefault("FOLLOWERS_SERVICE_PORT", "8001"),
		UsersHost:     LookupEnvOrGetDefault("USERS_SERVICE_HOST", "localhost"),
		UsersPort:     LookupEnvOrGetDefault("USERS_PORT", "9090"),
	}
}

func LookupEnvOrGetDefault(key string, defaultValue string) string {
	if env, found := os.LookupEnv(key); !found {
		return defaultValue
	} else {
		return env
	}
}
