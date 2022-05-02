package startup

import (
	"api_gateway/startup/config"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	followers "github.com/stojic19/XWS-TIM15/common/followers" //MORACE SA GITHUBA
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	config *config.Config
	mux    *runtime.ServeMux
}

func NewServer(config *config.Config) *Server {
	server := &Server{
		config: config,
		mux:    runtime.NewServeMux(),
	}
	return server
}

func (server *Server) InitHandlers() {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	followersEndpoint := fmt.Sprintf("%s:%s", server.config.FollowersHost, server.config.FollowersPort)
	err := followers.RegisterFollowersServiceHandlerFromEndpoint()
}
