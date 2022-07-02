package startup

import (
	"fmt"
	otgo "github.com/opentracing/opentracing-go"
	"github.com/stojic19/XWS-TIM15/common/proto/posts"
	"github.com/stojic19/XWS-TIM15/common/tracer"
	"github.com/stojic19/XWS-TIM15/posts_microservice/application"
	"github.com/stojic19/XWS-TIM15/posts_microservice/domain"
	"github.com/stojic19/XWS-TIM15/posts_microservice/infrastructure/api"
	"github.com/stojic19/XWS-TIM15/posts_microservice/infrastructure/persistence"
	"github.com/stojic19/XWS-TIM15/posts_microservice/startup/config"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

const (
	Name = "Posts service"
)

type Server struct {
	config *config.Config
	Tracer otgo.Tracer
	Closer io.Closer
}

func NewServer(config *config.Config) *Server {
	tracer, closer := tracer.Init(Name)
	otgo.SetGlobalTracer(tracer)
	return &Server{
		config: config,
		Tracer: tracer,
		Closer: closer,
	}
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.PostsDbHost, server.config.PostsDbPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	postsStore := server.initPostsStore(mongoClient)
	postsService := server.initPostsService(postsStore)
	postsHandler := server.initPostsHandler(postsService)
	server.startGrpcServer(postsHandler)
}

func (server *Server) initPostsStore(client *mongo.Client) domain.PostsStore {
	return persistence.NewPostsStore(client)
}

func (server *Server) initPostsService(store domain.PostsStore) *application.PostsService {
	return application.NewPostsService(store)
}

func (server *Server) initPostsHandler(postsService *application.PostsService) *api.PostsHandler {
	followersEndpoint := fmt.Sprintf("%s:%s", server.config.FollowersHost, server.config.FollowersPort)
	usersEndpoint := fmt.Sprintf("%s:%s", server.config.UsersHost, server.config.UsersPort)
	return api.NewPostsHandler(postsService, followersEndpoint, usersEndpoint)
}

func (server *Server) startGrpcServer(postsHandler *api.PostsHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	posts.RegisterPostsServiceServer(grpcServer, postsHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
