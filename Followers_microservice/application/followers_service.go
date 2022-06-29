package application

import (
	"github.com/stojic19/XWS-TIM15/Followers_microservice/domain"
)

type FollowersService struct {
	store               domain.FollowersStore
	orchestrator        *BlockOrchestrator
	unblockOrchestrator *UnblockOrchestrator
}

func NewFollowersService(store domain.FollowersStore, orchestrator *BlockOrchestrator, unblockOrchestrator *UnblockOrchestrator) *FollowersService {
	return &FollowersService{
		store:               store,
		orchestrator:        orchestrator,
		unblockOrchestrator: unblockOrchestrator,
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
	//return service.store.Block(blockerId, blockedId)
	response, err := service.store.BlockPending(blockerId, blockedId)
	if err != nil {
		return "Error", err
	}
	err = service.orchestrator.Start(blockerId, blockedId)
	if err != nil {
		service.store.RevertPendingBlock(blockerId, blockedId)
		return "Error during saga", err
	}
	return response, nil
}

func (service *FollowersService) ConfirmBlock(blockerId string, blockedId string) (string, error) {
	response, err := service.store.ConfirmBlock(blockerId, blockedId)
	if err != nil {
		return "Error", err
	}
	return response, nil
}

func (service *FollowersService) RevertBlock(blockerId string, blockedId string) (string, error) {
	response, err := service.store.RevertPendingBlock(blockerId, blockedId)
	if err != nil {
		return "Error", err
	}
	return response, nil
}

func (service *FollowersService) Unblock(blockerId string, blockedId string) (string, error) {
	response, err := service.store.UnblockPending(blockerId, blockedId)
	if err != nil {
		return "Error", err
	}
	if response == "failed to unblock, user is not blocked" {
		return response, nil
	}
	err = service.unblockOrchestrator.Start(blockerId, blockedId)
	if err != nil {
		service.store.RevertPendingUnblock(blockerId, blockedId)
		return "Error during saga", err
	}
	return response, nil
}

func (service *FollowersService) ConfirmUnblock(blockerId string, blockedId string) (string, error) {
	response, err := service.store.ConfirmUnblock(blockerId, blockedId)
	if err != nil {
		return "Error", err
	}
	return response, nil
}

func (service *FollowersService) RevertUnblock(blockerId string, blockedId string) (string, error) {
	response, err := service.store.RevertPendingUnblock(blockerId, blockedId)
	if err != nil {
		return "Error", err
	}
	return response, nil
}

func (service *FollowersService) GetBlocked(id string) ([]*domain.User, error) {
	return service.store.GetBlocked(id)
}
func (service *FollowersService) GetBlockers(id string) ([]*domain.User, error) {
	return service.store.GetBlockers(id)
}
func (service *FollowersService) GetRecommendedUsers(id string) ([]*domain.User, error) {
	return service.store.GetRecommended(id)
}
