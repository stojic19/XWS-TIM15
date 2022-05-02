package main

import (
	"github.com/stojic19/XWS-TIM15/Followers_microservice/startup"
	cfg "github.com/stojic19/XWS-TIM15/Followers_microservice/startup/config"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
