package startup

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	otgo "github.com/opentracing/opentracing-go"
	"github.com/stojic19/XWS-TIM15/Followers_microservice/application"
	"github.com/stojic19/XWS-TIM15/Followers_microservice/domain"
	"github.com/stojic19/XWS-TIM15/Followers_microservice/infrastructure/api"
	"github.com/stojic19/XWS-TIM15/Followers_microservice/infrastructure/persistence"
	"github.com/stojic19/XWS-TIM15/Followers_microservice/startup/config"
	"github.com/stojic19/XWS-TIM15/common/proto/followers"
	saga "github.com/stojic19/XWS-TIM15/common/saga/messaging"
	"github.com/stojic19/XWS-TIM15/common/saga/messaging/nats"
	"github.com/stojic19/XWS-TIM15/common/tracer"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"strings"
)

const (
	QueueGroup = "followers_service"
	Name       = "Followers service"
)

type Server struct {
	config *config.Config
	Tracer otgo.Tracer
	Closer io.Closer
}

type Neo4jConfiguration struct {
	Url      string
	Username string
	Password string
	Database string
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

func (server *Server) Start() {
	configuration := server.parseConfiguration()
	driver, err := configuration.NewDriver()
	if err != nil {
		log.Fatal(err)
	}
	followersStore := server.initFollowersStore(&driver, configuration.Database)

	commandPublisher := server.initPublisher(server.config.BlockCommandSubject)
	replySubscriber := server.initSubscriber(server.config.BlockReplySubject, QueueGroup)
	blockOrchestrator := server.initBlockOrchestrator(commandPublisher, replySubscriber)

	unblockCommandPublisher := server.initPublisher(server.config.UnblockCommandSubject)
	unblockReplySubscriber := server.initSubscriber(server.config.UnblockReplySubject, QueueGroup)
	unblockOrchestrator := server.initUnblockOrchestrator(unblockCommandPublisher, unblockReplySubscriber)

	followersService := server.initFollowersService(followersStore, blockOrchestrator, unblockOrchestrator)

	commandSubscriber := server.initSubscriber(server.config.BlockCommandSubject, QueueGroup)
	replyPublisher := server.initPublisher(server.config.BlockReplySubject)
	server.initBlockHandler(followersService, replyPublisher, commandSubscriber)

	unblockCommandSubscriber := server.initSubscriber(server.config.UnblockCommandSubject, QueueGroup)
	unblockReplyPublisher := server.initPublisher(server.config.UnblockReplySubject)
	server.initUnblockHandler(followersService, unblockReplyPublisher, unblockCommandSubscriber)

	followersHandler := server.initFollowersHandler(followersService)

	server.startGrpcServer(followersHandler)
}

func (server *Server) parseConfiguration() *Neo4jConfiguration {
	database := server.config.DbDatabase
	if !strings.HasPrefix(server.config.DbNeo4jVersion, "4") {
		database = ""
	}
	return &Neo4jConfiguration{
		Url: fmt.Sprintf("neo4j://%s:%s", server.config.DbHost, server.config.DbPort), //config.LookupEnvOrGetDefault("NEO4J_URI", "neo4j+s://demo.neo4jlabs.com")
		//Username: lookupEnvOrGetDefault("NEO4J_USER", "neo4j"),
		Username: server.config.DbUsername,
		//Password: lookupEnvOrGetDefault("NEO4J_PASSWORD", "password"),
		Password: server.config.DbPassword,
		Database: database,
	}
}

func (nc *Neo4jConfiguration) NewDriver() (neo4j.Driver, error) {
	return neo4j.NewDriver(nc.Url, neo4j.BasicAuth(nc.Username, nc.Password, ""))
}

func (server *Server) initFollowersStore(driver *neo4j.Driver, dbName string) domain.FollowersStore {
	return persistence.NewFollowersStore(driver, dbName)
}

func (server *Server) initPublisher(subject string) saga.Publisher {
	publisher, err := nats.NewNATSPublisher(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject)
	if err != nil {
		log.Fatal(err)
	}
	return publisher
}

func (server *Server) initSubscriber(subject, queueGroup string) saga.Subscriber {
	subscriber, err := nats.NewNATSSubscriber(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject, queueGroup)
	if err != nil {
		log.Fatal(err)
	}
	return subscriber
}

func (server *Server) initBlockOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) *application.BlockOrchestrator {
	orchestrator, err := application.NewBlockOrchestrator(publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
	return orchestrator
}

func (server *Server) initUnblockOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) *application.UnblockOrchestrator {
	orchestrator, err := application.NewUnblockOrchestrator(publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
	return orchestrator
}

func (server *Server) initFollowersService(store domain.FollowersStore, orchestrator *application.BlockOrchestrator, unblockOrchestrator *application.UnblockOrchestrator) *application.FollowersService {
	return application.NewFollowersService(store, orchestrator, unblockOrchestrator)
}

func (server *Server) initBlockHandler(service *application.FollowersService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := api.NewBlockCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initUnblockHandler(service *application.FollowersService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := api.NewUnblockCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initFollowersHandler(followersService *application.FollowersService) *api.FollowersHandler {
	usersEndpoint := fmt.Sprintf("%s:%s", server.config.UsersHost, server.config.UsersPort)
	return api.NewFollowersHandler(followersService, usersEndpoint)
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
