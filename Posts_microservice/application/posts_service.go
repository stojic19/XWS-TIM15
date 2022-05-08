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

func (service *PostsService) LikePost(postId primitive.ObjectID, user *domain.User) error {
	err := service.store.RemoveDislike(postId, user)
	if err != nil {
		return err
	}
	return service.store.LikePost(postId, user)
}

func (service *PostsService) RemoveLike(postId primitive.ObjectID, user *domain.User) error {
	return service.store.RemoveLike(postId, user)
}

func (service *PostsService) DislikePost(postId primitive.ObjectID, user *domain.User) error {
	err := service.store.RemoveLike(postId, user)
	if err != nil {
		return err
	}
	return service.store.DislikePost(postId, user)
}

func (service *PostsService) RemoveDislike(postId primitive.ObjectID, user *domain.User) error {
	return service.store.RemoveDislike(postId, user)
}

func (service *PostsService) CreateComment(postId primitive.ObjectID, comment *domain.Comment) error {
	return service.store.CreateComment(postId, comment)
}
