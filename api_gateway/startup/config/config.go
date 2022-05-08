package config

import "os"

type Config struct {
	Port          string
	FollowersHost string
	FollowersPort string
	UsersHost     string
	UsersPort     string
}

func NewConfig() *Config {
	return &Config{
		Port:          os.Getenv("GATEWAY_PORT"),
		FollowersHost: os.Getenv("FOLLOWERS_SERVICE_HOST"),
		FollowersPort: os.Getenv("FOLLOWERS_SERVICE_PORT"),
		UsersHost:     os.Getenv("USERS_SERVICE_HOST"),
		UsersPort:     os.Getenv("USERS_PORT"),
	}
}

func LookupEnvOrGetDefault(key string, defaultValue string) string {
	if env, found := os.LookupEnv(key); !found {
		return defaultValue
	} else {
		return env
	}
}
