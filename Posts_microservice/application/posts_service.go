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

func (service *PostsService) GetAll() ([]*domain.Post, error) {
	return service.store.GetAll()
}

func (service *PostsService) CreatePost(post *domain.Post) error {
	return service.store.Create(post)
}
