package startup

import (
	"context"
	"fmt"
	otgo "github.com/opentracing/opentracing-go"
	"github.com/rs/cors"
	"github.com/stojic19/XWS-TIM15/api_gateway/infrastructure/api"
	"github.com/stojic19/XWS-TIM15/api_gateway/startup/config"
	"github.com/stojic19/XWS-TIM15/api_gateway/startup/middleware"
	"github.com/stojic19/XWS-TIM15/common/proto/chat"
	"github.com/stojic19/XWS-TIM15/common/proto/followers"
	"github.com/stojic19/XWS-TIM15/common/proto/job_offers"
	"github.com/stojic19/XWS-TIM15/common/proto/notifications"
	"github.com/stojic19/XWS-TIM15/common/proto/posts"
	"github.com/stojic19/XWS-TIM15/common/proto/users"
	"github.com/stojic19/XWS-TIM15/common/tracer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"net/http"
)

type Server struct {
	config *config.Config
	Tracer otgo.Tracer
	Closer io.Closer
	mux    *MuxWrapper
}

const (
	Name = "Api-composition"
)

func NewServer(config *config.Config) *Server {
	tracer, closer := tracer.Init(Name)
	otgo.SetGlobalTracer(tracer)
	server := &Server{
		config: config,
		mux:    NewMuxWrapper(),
		Tracer: tracer,
		Closer: closer,
	}
	server.initHandlers()
	server.initCustomHandlers()
	return server
}

func (server *Server) initHandlers() {
	server.mux.AppendMiddleware(middleware.AuthMiddleware)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	fmt.Printf("Followers: %s:%s\n", server.config.FollowersHost, server.config.FollowersPort)
	followersEndpoint := fmt.Sprintf("%s:%s", server.config.FollowersHost, server.config.FollowersPort)
	err := followers.RegisterFollowersServiceHandlerFromEndpoint(context.TODO(), &server.mux.ServeMux, followersEndpoint, opts)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Posts: %s:%s\n", server.config.PostsHost, server.config.PostsPort)
	postsEndpoint := fmt.Sprintf("%s:%s", server.config.PostsHost, server.config.PostsPort)
	err = posts.RegisterPostsServiceHandlerFromEndpoint(context.TODO(), &server.mux.ServeMux, postsEndpoint, opts)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Users: %s:%s\n", server.config.UsersHost, server.config.UsersPort)
	usersEndpoint := fmt.Sprintf("%s:%s", server.config.UsersHost, server.config.UsersPort)
	err = users.RegisterUsersServiceHandlerFromEndpoint(context.TODO(), &server.mux.ServeMux, usersEndpoint, opts)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Job offers: %s:%s\n", server.config.JobOffersHost, server.config.JobOffersPort)
	jobOffersEndpoint := fmt.Sprintf("%s:%s", server.config.JobOffersHost, server.config.JobOffersPort)
	err = job_offers.RegisterJobOffersServiceHandlerFromEndpoint(context.TODO(), &server.mux.ServeMux, jobOffersEndpoint, opts)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Chat: %s:%s\n", server.config.ChatHost, server.config.ChatPort)
	chatEndpoint := fmt.Sprintf("%s:%s", server.config.ChatHost, server.config.ChatPort)
	err = chat.RegisterChatServiceGrpcHandlerFromEndpoint(context.TODO(), &server.mux.ServeMux, chatEndpoint, opts)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Notifications: %s:%s\n", server.config.NotificationsHost, server.config.NotificationsPort)
	notificationsEndpoint := fmt.Sprintf("%s:%s", server.config.NotificationsHost, server.config.NotificationsPort)
	err = notifications.RegisterNotificationsServiceHandlerFromEndpoint(context.TODO(), &server.mux.ServeMux, notificationsEndpoint, opts)
	if err != nil {
		panic(err)
	}
}

func (server *Server) initCustomHandlers() {
	followersEndpoint := fmt.Sprintf("%s:%s", server.config.FollowersHost, server.config.FollowersPort)
	usersEndpoint := fmt.Sprintf("%s:%s", server.config.UsersHost, server.config.UsersPort)
	postsEndpoint := fmt.Sprintf("%s:%s", server.config.PostsHost, server.config.PostsPort)
	jobOffersEndpoint := fmt.Sprintf("%s:%s", server.config.JobOffersHost, server.config.JobOffersPort)
	followersHandler := api.NewFollowersHandler(followersEndpoint, usersEndpoint)
	followersHandler.Init(&server.mux.ServeMux)
	postsHandler := api.NewPostsHandler(postsEndpoint, followersEndpoint, usersEndpoint)
	postsHandler.Init(&server.mux.ServeMux)
	jobOffersHandler := api.NewJobOffersHandler(jobOffersEndpoint, usersEndpoint)
	jobOffersHandler.Init(&server.mux.ServeMux)
}

func (server *Server) Start() {
	fmt.Printf("Port: %s\n", server.config.Port)
	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "https://localhost:3000/**", "http://localhost:3002", "https://localhost:3002/**"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Accept-Language", "Content-Type", "Content-Language", "Origin", "Authorization", "Access-Control-Allow-Origin", "*"},
		AllowCredentials: true,
	}).Handler(server.mux)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", server.config.Port), handler))
}
