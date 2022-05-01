package main

import (
	"Followers_microservice/startup"
	cfg "Followers_microservice/startup/config"
)

func main(){
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}