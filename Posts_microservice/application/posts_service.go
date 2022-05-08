package application

import (
	"github.com/stojic19/XWS-TIM15/posts_microservice/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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

func (service *PostsService) Get(id primitive.ObjectID) (*domain.Post, error) {
	return service.store.Get(id)
}

func (service *PostsService) GetFromUser(id string) ([]*domain.Post, error) {
	return service.store.GetFromUser(id)
}

func (service *PostsService) CreatePost(post *domain.Post) error {
	return service.store.Create(post)
}
