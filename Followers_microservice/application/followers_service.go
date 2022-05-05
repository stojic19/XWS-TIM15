package application

import (
	"github.com/stojic19/XWS-TIM15/Followers_microservice/domain"
	"strings"
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
func (service *FollowersService) Follow(followerUsername string, followedUsername string) (string, error) {
	//kad se napravi profile service, ovde se pita da li je profil privatan
	if strings.HasPrefix(followedUsername, "p") {
		return service.store.FollowRequest(followerUsername, followedUsername)
	}
	return service.store.Follow(followerUsername, followedUsername)
}
