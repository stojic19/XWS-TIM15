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
func (service *FollowersService) GetFollows(id string) ([]*domain.User, error) {
	return service.store.GetFollows(id)
}
func (service *FollowersService) GetFollowers(id string) ([]*domain.User, error) {
	return service.store.GetFollowers(id)
}
func (service *FollowersService) GetFollowRequests(id string) ([]*domain.User, error) {
	return service.store.GetFollowRequests(id)
}
func (service *FollowersService) GetFollowerRequests(id string) ([]*domain.User, error) {
	return service.store.GetFollowerRequests(id)
}
func (service *FollowersService) GetRelationship(followerId string, followedId string) (string, error) {
	return service.store.GetRelationship(followerId, followedId)
}
func (service *FollowersService) Follow(followerId string, followedId string) (string, error) {
	return service.store.Follow(followerId, followedId)
}
func (service *FollowersService) FollowRequest(followerId string, followedId string) (string, error) {
	return service.store.FollowRequest(followerId, followedId)
}
func (service *FollowersService) ConfirmFollow(followerId string, followedId string) (string, error) {
	return service.store.ConfirmFollow(followerId, followedId)
}
func (service *FollowersService) Unfollow(followerId string, followedId string) (string, error) {
	return service.store.Unfollow(followerId, followedId)
}
func (service *FollowersService) RemoveFollowRequest(followerId string, followedId string) (string, error) {
	return service.store.RemoveFollowRequest(followerId, followedId)
}
func (service *FollowersService) Block(blockerId string, blockedId string) (string, error) {
	return service.store.Block(blockerId, blockedId)
}
func (service *FollowersService) Unblock(blockerId string, blockedId string) (string, error) {
	return service.store.Unblock(blockerId, blockedId)
}
