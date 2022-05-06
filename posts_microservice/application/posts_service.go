package application

import "github.com/stojic19/XWS-TIM15/posts_microservice/domain"

type PostsService struct {
	store domain.PostsStore
}

func NewPostsService(store domain.PostsStore) *PostsService {
	return &PostsService{
		store: store,
	}
}
