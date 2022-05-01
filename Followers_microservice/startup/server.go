package startup

import (
	"Followers_microservice/infrastructure/api"
	followers "Followers_microservice/proto"
	"Followers_microservice/startup/config"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (server *Server) Start() {
	followersHandler := server.initFollowersHandler()
	server.startGrpcServer(followersHandler)
}

func (server *Server) initFollowersHandler() *api.FollowersHandler {
	return api.NewFollowersHandler()
}

func (server *Server) startGrpcServer(followersHandler *api.FollowersHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	followers.RegisterFollowersServiceServer(grpcServer, followersHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
