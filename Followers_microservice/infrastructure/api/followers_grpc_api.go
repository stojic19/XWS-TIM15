package api

import (
	"context"
	"github.com/stojic19/XWS-TIM15/Followers_microservice/application"
	"github.com/stojic19/XWS-TIM15/common/proto/followers"
)

type FollowersHandler struct {
	followers.UnimplementedFollowersServiceServer
	service *application.FollowersService
}

func NewFollowersHandler(service *application.FollowersService) *FollowersHandler {
	return &FollowersHandler{
		service: service,
	}
}

func (handler *FollowersHandler) GetFollowing(ctx context.Context, request *followers.GetFollowingRequest) (*followers.GetFollowingResponse, error) {
	username := request.Username
	response, err := handler.service.GetFollowing(username)
	if err != nil {
		return nil, err
	}
	responsePb := &followers.GetFollowingResponse{Followers: []*followers.Follower{}}
	for _, user := range response {
		responsePb.Followers = append(responsePb.Followers, &followers.Follower{Username: user.Username})
	}
	return responsePb, nil
}

func (handler *FollowersHandler) ConfirmFollow(ctx context.Context, request *followers.ConfirmFollowRequest) (*followers.ConfirmFollowResponse, error) {
	return nil, nil
}

func (handler *FollowersHandler) Follow(ctx context.Context, request *followers.FollowRequest) (*followers.FollowResponse, error) {
	return nil, nil
}
