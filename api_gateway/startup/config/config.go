package config

import "os"

type Config struct {
	Port          string
	FollowersHost string
	FollowersPort string
	PostsHost     string
	PostsPort     string
}

func NewConfig() *Config {
	return &Config{
		Port:          os.Getenv("GATEWAY_PORT"),
		FollowersHost: os.Getenv("FOLLOWERS_SERVICE_HOST"),
		FollowersPort: os.Getenv("FOLLOWERS_SERVICE_PORT"),
		PostsHost:     os.Getenv("POSTS_SERVICE_HOST"),
		PostsPort:     os.Getenv("POSTS_SERVICE_PORT"),
	}
}
