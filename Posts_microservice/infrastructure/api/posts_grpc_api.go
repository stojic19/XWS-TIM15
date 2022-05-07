package api

import "github.com/stojic19/XWS-TIM15/posts_microservice/application"

type PostsHandler struct {
	//unimplementedservice
	service *application.PostsService
}

func NewPostsHandler(service *application.PostsService) *PostsHandler {
	return &PostsHandler{
		service: service,
	}
}
