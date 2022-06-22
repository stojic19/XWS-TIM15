package startup

import (
	"fmt"
	"github.com/stojic19/XWS-TIM15/common/proto/job_offers"
	"github.com/stojic19/XWS-TIM15/job_offers_microservice/application"
	"github.com/stojic19/XWS-TIM15/job_offers_microservice/domain"
	"github.com/stojic19/XWS-TIM15/job_offers_microservice/infrastructure/api"
	"github.com/stojic19/XWS-TIM15/job_offers_microservice/infrastructure/persistence"
	"github.com/stojic19/XWS-TIM15/job_offers_microservice/startup/config"
	"go.mongodb.org/mongo-driver/mongo"
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

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.JobOffersDbHost, server.config.JobOffersDbPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	jobOffersStore := server.initJobOffersStore(mongoClient)
	jobOffersService := server.initJobOffersService(jobOffersStore)
	jobOffersHandler := server.initJobOffersHandler(jobOffersService)
	server.startGrpcServer(jobOffersHandler)
}

func (server *Server) initJobOffersStore(client *mongo.Client) domain.JobOffersStore {
	return persistence.NewJobOffersStore(client)
}

func (server *Server) initJobOffersService(store domain.JobOffersStore) *application.JobOffersService {
	return application.NewJobOffersService(store)
}

func (server *Server) initJobOffersHandler(jobOffersService *application.JobOffersService) *api.JobOffersHandler {
	usersEndpoint := fmt.Sprintf("%s:%s", server.config.UsersHost, server.config.UsersPort)
	return api.NewJobOffersHandler(jobOffersService, usersEndpoint)
}

func (server *Server) startGrpcServer(jobOffersHandler *api.JobOffersHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	job_offers.RegisterJobOffersServiceServer(grpcServer, jobOffersHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
