package config

import "os"

type Config struct {
	Port              string
	FollowersHost     string
	FollowersPort     string
	PostsHost         string
	PostsPort         string
	UsersHost         string
	UsersPort         string
	JobOffersHost     string
	JobOffersPort     string
	ChatHost          string
	ChatPort          string
	NotificationsHost string
	NotificationsPort string
}

func NewConfig() *Config {
	return &Config{
		Port:              os.Getenv("GATEWAY_PORT"),
		FollowersHost:     os.Getenv("FOLLOWERS_SERVICE_HOST"),
		FollowersPort:     os.Getenv("FOLLOWERS_SERVICE_PORT"),
		PostsHost:         os.Getenv("POSTS_SERVICE_HOST"),
		PostsPort:         os.Getenv("POSTS_SERVICE_PORT"),
		UsersHost:         os.Getenv("USERS_SERVICE_HOST"),
		UsersPort:         os.Getenv("USERS_PORT"),
		JobOffersPort:     LookupEnvOrGetDefault("JOB_OFFERS_SERVICE_PORT", "8003"),
		JobOffersHost:     LookupEnvOrGetDefault("JOB_OFFERS_SERVICE_HOST", "localhost"),
		ChatHost:          LookupEnvOrGetDefault("CHAT_HOST", "localhost"),
		ChatPort:          LookupEnvOrGetDefault("CHAT_PORT", "8004"),
		NotificationsHost: LookupEnvOrGetDefault("NOTIFICATIONS_HOST", "localhost"),
		NotificationsPort: LookupEnvOrGetDefault("NOTIFICATIONS_PORT", "8006"),
	}
}

func LookupEnvOrGetDefault(key string, defaultValue string) string {
	if env, found := os.LookupEnv(key); !found {
		return defaultValue
	} else {
		return env
	}
}
