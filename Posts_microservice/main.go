package main

import (
	"github.com/stojic19/XWS-TIM15/posts_microservice/startup"
	"github.com/stojic19/XWS-TIM15/posts_microservice/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
