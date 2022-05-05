package application

import (
	"github.com/stojic19/XWS-TIM15/Followers_microservice/domain"
)

type FollowersService struct {
	store domain.FollowersStore
}

func NewFollowersService(store domain.FollowersStore) *FollowersService {
	return &FollowersService{
		store: store,
	}
}
func (service *FollowersService) GetFollowing(username string) ([]*domain.User, error) {
	return service.store.GetFollowing(username)
}
