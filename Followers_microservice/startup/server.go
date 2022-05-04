package startup

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/stojic19/XWS-TIM15/Followers_microservice/application"
	"github.com/stojic19/XWS-TIM15/Followers_microservice/domain"
	"github.com/stojic19/XWS-TIM15/Followers_microservice/infrastructure/api"
	"github.com/stojic19/XWS-TIM15/Followers_microservice/infrastructure/persistence"
	"github.com/stojic19/XWS-TIM15/Followers_microservice/startup/config"
	"github.com/stojic19/XWS-TIM15/common/proto/followers"
	"google.golang.org/grpc"
	"log"
	"net"
	"strings"
)

type Server struct {
	config *config.Config
}

type Neo4jConfiguration struct {
	Url      string
	Username string
	Password string
	Database string
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (server *Server) Start() {
	configuration := server.parseConfiguration()
	driver, err := configuration.NewDriver()
	if err != nil {
		log.Fatal(err)
	}
	followersStore := server.initFollowersStore(&driver, configuration.Database)

	followersService := server.initFollowersService(followersStore)
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

func (server *Server) initFollowersService(store domain.FollowersStore) *application.FollowersService {
	return application.NewFollowersService(store)
}

func (server *Server) initFollowersHandler(followersService *application.FollowersService) *api.FollowersHandler {
	return api.NewFollowersHandler(followersService)
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
