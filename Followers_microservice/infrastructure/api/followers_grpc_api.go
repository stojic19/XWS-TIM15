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

func (handler *FollowersHandler) GetFollows(ctx context.Context, request *followers.GetFollowsRequest) (*followers.GetFollowsResponse, error) {
	username := request.Username
	response, err := handler.service.GetFollows(username)
	if err != nil {
		return nil, err
	}
	responsePb := &followers.GetFollowsResponse{Followers: []*followers.Follower{}}
	for _, user := range response {
		responsePb.Followers = append(responsePb.Followers, &followers.Follower{Username: user.Username})
	}
	return responsePb, nil
}

func (handler *FollowersHandler) GetFollowers(ctx context.Context, request *followers.GetFollowersRequest) (*followers.GetFollowersResponse, error) {
	username := request.Username
	response, err := handler.service.GetFollowers(username)
	if err != nil {
		return nil, err
	}
	responsePb := &followers.GetFollowersResponse{Followers: []*followers.Follower{}}
	for _, user := range response {
		responsePb.Followers = append(responsePb.Followers, &followers.Follower{Username: user.Username})
	}
	return responsePb, nil
}

func (handler *FollowersHandler) ConfirmFollow(ctx context.Context, request *followers.ConfirmFollowRequest) (*followers.ConfirmFollowResponse, error) {
	followerUsername := request.FollowerUsername
	followedUsername := request.FollowedUsername
	response, err := handler.service.ConfirmFollow(followerUsername, followedUsername)
	if err != nil {
		return nil, err
	}
	responsePb := &followers.ConfirmFollowResponse{Response: response}
	return responsePb, nil
}

func (handler *FollowersHandler) Follow(ctx context.Context, request *followers.FollowRequest) (*followers.FollowResponse, error) {
	followerUsername := request.FollowerUsername
	followedUsername := request.FollowedUsername
	response, err := handler.service.Follow(followerUsername, followedUsername)
	if err != nil {
		return nil, err
	}
	responsePb := &followers.FollowResponse{Response: response}
	return responsePb, nil
}
