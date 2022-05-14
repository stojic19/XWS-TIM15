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
	id := request.Id
	response, err := handler.service.GetFollows(id)
	if err != nil {
		return nil, err
	}
	responsePb := &followers.GetFollowsResponse{Follows: []*followers.Follower{}}
	for _, user := range response {
		responsePb.Follows = append(responsePb.Follows, &followers.Follower{Id: user.Id})
	}
	return responsePb, nil
}

func (handler *FollowersHandler) GetFollowers(ctx context.Context, request *followers.GetFollowersRequest) (*followers.GetFollowersResponse, error) {
	id := request.Id
	response, err := handler.service.GetFollowers(id)
	if err != nil {
		return nil, err
	}
	responsePb := &followers.GetFollowersResponse{Followers: []*followers.Follower{}}
	for _, user := range response {
		responsePb.Followers = append(responsePb.Followers, &followers.Follower{Id: user.Id})
	}
	return responsePb, nil
}

func (handler *FollowersHandler) GetFollowRequests(ctx context.Context, request *followers.GetFollowRequestsRequest) (*followers.GetFollowRequestsResponse, error) {
	id := request.Id
	response, err := handler.service.GetFollowRequests(id)
	if err != nil {
		return nil, err
	}
	responsePb := &followers.GetFollowRequestsResponse{FollowRequests: []*followers.Follower{}}
	for _, user := range response {
		responsePb.FollowRequests = append(responsePb.FollowRequests, &followers.Follower{Id: user.Id})
	}
	return responsePb, nil
}

func (handler *FollowersHandler) GetFollowerRequests(ctx context.Context, request *followers.GetFollowerRequestsRequest) (*followers.GetFollowerRequestsResponse, error) {
	id := request.Id
	response, err := handler.service.GetFollowerRequests(id)
	if err != nil {
		return nil, err
	}
	responsePb := &followers.GetFollowerRequestsResponse{FollowerRequests: []*followers.Follower{}}
	for _, user := range response {
		responsePb.FollowerRequests = append(responsePb.FollowerRequests, &followers.Follower{Id: user.Id})
	}
	return responsePb, nil
}

func (handler *FollowersHandler) ConfirmFollow(ctx context.Context, request *followers.ConfirmFollowRequest) (*followers.ConfirmFollowResponse, error) {
	followerId := request.FollowerId
	followedId := request.FollowedId
	response, err := handler.service.ConfirmFollow(followerId, followedId)
	if err != nil {
		return nil, err
	}
	responsePb := &followers.ConfirmFollowResponse{Response: response}
	return responsePb, nil
}

func (handler *FollowersHandler) Follow(ctx context.Context, request *followers.FollowRequest) (*followers.FollowResponse, error) {
	followerId := request.FollowerId
	followedId := request.FollowedId
	response, err := handler.service.Follow(followerId, followedId)
	if err != nil {
		return nil, err
	}
	responsePb := &followers.FollowResponse{Response: response}
	return responsePb, nil
}

func (handler *FollowersHandler) Unfollow(ctx context.Context, request *followers.UnfollowRequest) (*followers.UnfollowResponse, error) {
	followerId := request.FollowerId
	followedId := request.FollowedId
	response, err := handler.service.Unfollow(followerId, followedId)
	if err != nil {
		return nil, err
	}
	responsePb := &followers.UnfollowResponse{Response: response}
	return responsePb, nil
}

func (handler *FollowersHandler) RemoveFollowRequest(ctx context.Context, request *followers.RemoveFollowRequestRequest) (*followers.RemoveFollowRequestResponse, error) {
	followerId := request.FollowerId
	followedId := request.FollowedId
	response, err := handler.service.RemoveFollowRequest(followerId, followedId)
	if err != nil {
		return nil, err
	}
	responsePb := &followers.RemoveFollowRequestResponse{Response: response}
	return responsePb, nil
}
