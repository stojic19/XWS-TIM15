package startup

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/stojic19/XWS-TIM15/api_gateway/infrastructure/api"
	"github.com/stojic19/XWS-TIM15/api_gateway/startup/config"
	"github.com/stojic19/XWS-TIM15/common/proto/followers"
	"github.com/stojic19/XWS-TIM15/common/proto/job_offers"
	"github.com/stojic19/XWS-TIM15/common/proto/posts"
	"github.com/stojic19/XWS-TIM15/common/proto/users"
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
	server.initCustomHandlers()
	return server
}

func (server *Server) initHandlers() {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	fmt.Printf("%s:%s\n", server.config.FollowersHost, server.config.FollowersPort)
	followersEndpoint := fmt.Sprintf("%s:%s", server.config.FollowersHost, server.config.FollowersPort)
	err := followers.RegisterFollowersServiceHandlerFromEndpoint(context.TODO(), server.mux, followersEndpoint, opts)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s:%s\n", server.config.PostsHost, server.config.PostsPort)
	postsEndpoint := fmt.Sprintf("%s:%s", server.config.PostsHost, server.config.PostsPort)
	err = posts.RegisterPostsServiceHandlerFromEndpoint(context.TODO(), server.mux, postsEndpoint, opts)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s:%s\n", server.config.UsersHost, server.config.UsersPort)
	usersEndpoint := fmt.Sprintf("%s:%s", server.config.UsersHost, server.config.UsersPort)
	err = users.RegisterUsersServiceHandlerFromEndpoint(context.TODO(), server.mux, usersEndpoint, opts)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s:%s", server.config.JobOffersHost, server.config.JobOffersPort)
	jobOffersEndpoint := fmt.Sprintf("%s:%s", server.config.JobOffersHost, server.config.JobOffersPort)
	err = job_offers.RegisterJobOffersServiceHandlerFromEndpoint(context.TODO(), server.mux, jobOffersEndpoint, opts)
	if err != nil {
		panic(err)
	}
}

func (server *Server) initCustomHandlers() {
	followersEndpoint := fmt.Sprintf("%s:%s", server.config.FollowersHost, server.config.FollowersPort)
	usersEndpoint := fmt.Sprintf("%s:%s", server.config.UsersHost, server.config.UsersPort)
	postsEndpoint := fmt.Sprintf("%s:%s", server.config.PostsHost, server.config.PostsPort)
	followersHandler := api.NewFollowersHandler(followersEndpoint, usersEndpoint)
	followersHandler.Init(server.mux)
	postsHandler := api.NewPostsHandler(postsEndpoint, followersEndpoint, usersEndpoint)
	postsHandler.Init(server.mux)
}

func (server *Server) Start() {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), server.mux))
}
