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
	"os"
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
	configuration := parseConfiguration()
	driver, err := configuration.NewDriver()
	if err != nil {
		log.Fatal(err)
	}
	followersStore := server.initFollowersStore(&driver)

	followersService := server.initFollowersService(followersStore)
	followersHandler := server.initFollowersHandler(followersService)

	server.startGrpcServer(followersHandler)
}

func parseConfiguration() *Neo4jConfiguration {
	database := lookupEnvOrGetDefault("NEO4J_DATABASE", "followers")
	if !strings.HasPrefix(lookupEnvOrGetDefault("NEO4J_VERSION", "4"), "4") {
		database = ""
	}
	return &Neo4jConfiguration{
		Url: lookupEnvOrGetDefault("NEO4J_URI", "neo4j+s://demo.neo4jlabs.com"),
		//Username: lookupEnvOrGetDefault("NEO4J_USER", "neo4j"),
		Username: "neo4j",
		//Password: lookupEnvOrGetDefault("NEO4J_PASSWORD", "password"),
		Password: "password",
		Database: database,
	}
}

func lookupEnvOrGetDefault(key string, defaultValue string) string {
	if env, found := os.LookupEnv(key); !found {
		return defaultValue
	} else {
		return env
	}
}

func (nc *Neo4jConfiguration) NewDriver() (neo4j.Driver, error) {
	return neo4j.NewDriver(nc.Url, neo4j.BasicAuth(nc.Username, nc.Password, ""))
}

func (server *Server) initFollowersStore(driver *neo4j.Driver) domain.FollowersStore {
	return persistence.NewFollowersStore(driver)
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
