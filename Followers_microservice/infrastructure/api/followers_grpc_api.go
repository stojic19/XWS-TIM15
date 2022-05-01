package api

import (
	followers "Followers_microservice/proto" //MORACE SA GITHUBA
	"context"
)

type FollowersHandler struct {
	followers.UnimplementedFollowersServiceServer
}

func NewFollowersHandler() *FollowersHandler {
	return &FollowersHandler{}
}

func (handler *FollowersHandler) GetFollowing(ctx context.Context, request *followers.GetFollowingRequest) (*followers.GetFollowingResponse, error) {
	username := request.Username
	hardcodedUsernames := []string{username}
	hardcodedUsernames = append(hardcodedUsernames, "prvi")
	hardcodedUsernames = append(hardcodedUsernames, "drugi")
	response := &followers.GetFollowingResponse{
		Followers: []*followers.Follower{},
	}
	for _, user := range hardcodedUsernames {
		response.Followers = append(response.Followers, &followers.Follower{
			Username: user,
		})
	}
	return response, nil
}

func (handler *FollowersHandler) ConfirmFollow(ctx context.Context, request *followers.ConfirmFollowRequest) (*followers.ConfirmFollowResponse, error) {
	return nil, nil
}

func (handler *FollowersHandler) Follow(ctx context.Context, request *followers.FollowRequest) (*followers.FollowResponse, error) {
	return nil, nil
}
