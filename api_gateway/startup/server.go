package startup

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/stojic19/XWS-TIM15/api_gateway/startup/config"
	"github.com/stojic19/XWS-TIM15/common/proto/followers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
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
	server.initHandlers()
	return server
}

func (server *Server) initHandlers() {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	followersEndpoint := fmt.Sprintf("%s:%s", server.config.FollowersHost, server.config.FollowersPort)
	err := followers.RegisterFollowersServiceHandlerFromEndpoint(context.TODO(), server.mux, followersEndpoint, opts)
	if err != nil {
		panic(err)
	}
}

func (server *Server) Start() {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), server.mux))
}
