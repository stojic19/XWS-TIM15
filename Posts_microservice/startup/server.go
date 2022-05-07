package startup

import (
	"fmt"
	"github.com/stojic19/XWS-TIM15/posts_microservice/application"
	"github.com/stojic19/XWS-TIM15/posts_microservice/domain"
	"github.com/stojic19/XWS-TIM15/posts_microservice/infrastructure/api"
	"github.com/stojic19/XWS-TIM15/posts_microservice/infrastructure/persistence"
	"github.com/stojic19/XWS-TIM15/posts_microservice/startup/config"
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
	postsStore := server.initPostsStore()
	postsService := server.initPostsService(postsStore)
	postsHandler := server.initPostsHandler(postsService)
	server.startGrpcServer(postsHandler)
}

func (server *Server) initPostsStore() domain.PostsStore {
	return persistence.NewPostsStore()
}

func (server *Server) initPostsService(store domain.PostsStore) *application.PostsService {
	return application.NewPostsService(store)
}

func (server *Server) initPostsHandler(postsService *application.PostsService) *api.PostsHandler {
	return api.NewPostsHandler(postsService)
}

func (server *Server) startGrpcServer(postsHandler *api.PostsHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	//posts.RegisterPostsServiceServer(grpcServer, postsHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
