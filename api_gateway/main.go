package main

import (
	"github.com/stojic19/XWS-TIM15/api_gateway/startup"
	"github.com/stojic19/XWS-TIM15/api_gateway/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
