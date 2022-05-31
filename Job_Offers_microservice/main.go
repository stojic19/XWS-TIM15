package main

import (
	"github.com/stojic19/XWS-TIM15/job_offers_microservice/startup"
	"github.com/stojic19/XWS-TIM15/job_offers_microservice/startup/config"
)

func main() {
	config := config.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
