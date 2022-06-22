package api

import (
	"context"
	"fmt"
	"github.com/stojic19/XWS-TIM15/Followers_microservice/application"
	"github.com/stojic19/XWS-TIM15/Followers_microservice/infrastructure/services"
	"github.com/stojic19/XWS-TIM15/common/proto/followers"
	"github.com/stojic19/XWS-TIM15/common/proto/users"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FollowersHandler struct {
	followers.UnimplementedFollowersServiceServer
	service            *application.FollowersService
	usersClientAddress string
}

func NewFollowersHandler(service *application.FollowersService, usersEndpoint string) *FollowersHandler {
	return &FollowersHandler{
		service:            service,
		usersClientAddress: usersEndpoint,
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
		responsePb.Follows = append(responsePb.Follows, &followers.Follower{Id: user.Id, Time: timestamppb.New(user.TimeOfFollow)})
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
		responsePb.Followers = append(responsePb.Followers, &followers.Follower{Id: user.Id, Time: timestamppb.New(user.TimeOfFollow)})
	}
	return responsePb, nil
}

func (handler *FollowersHandler) GetFollowRequests(ctx context.Context, request *followers.GetFollowRequestsRequest) (*followers.GetFollowRequestsResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	fmt.Println(md)
	id := request.Id
	response, err := handler.service.GetFollowRequests(id)
	if err != nil {
		return nil, err
	}
	responsePb := &followers.GetFollowRequestsResponse{FollowRequests: []*followers.Follower{}}
	for _, user := range response {
		responsePb.FollowRequests = append(responsePb.FollowRequests, &followers.Follower{Id: user.Id, Time: timestamppb.New(user.TimeOfFollow)})
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
		responsePb.FollowerRequests = append(responsePb.FollowerRequests, &followers.Follower{Id: user.Id, Time: timestamppb.New(user.TimeOfFollow)})
	}
	return responsePb, nil
}

func (handler *FollowersHandler) GetRelationship(ctx context.Context, request *followers.GetRelationshipRequest) (*followers.GetRelationshipResponse, error) {
	followerId := request.FollowerId
	followedId := request.FollowedId
	response, err := handler.service.GetRelationship(followerId, followedId)
	if err != nil {
		return nil, err
	}
	responsePb := &followers.GetRelationshipResponse{}
	responsePb.Relationship = response
	return responsePb, nil
}

func (handler *FollowersHandler) ConfirmFollow(ctx context.Context, request *followers.ConfirmFollowRequest) (*followers.ConfirmFollowResponse, error) {
	//Endpoint protection
	metadata, _ := metadata.FromIncomingContext(ctx)
	sub := metadata.Get("sub")
	if sub == nil || sub[0] == "" {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized")
	}
	//Endpoint protection
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
	//Endpoint protection
	metadata, _ := metadata.FromIncomingContext(ctx)
	sub := metadata.Get("sub")
	if sub == nil || sub[0] == "" {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized")
	}
	//Endpoint protection
	followerId := request.FollowerId
	followedId := request.FollowedId
	userClient := services.NewUsersClient(handler.usersClientAddress)
	userResponse, err := userClient.GetUser(context.TODO(), &users.GetUserRequest{Id: followedId})
	if err != nil {
		return nil, err
	}
	if userResponse.User.IsPrivate {
		response, err := handler.service.FollowRequest(followerId, followedId)
		if err != nil {
			return nil, err
		}
		responsePb := &followers.FollowResponse{Response: response}
		return responsePb, nil
	}
	response, err := handler.service.Follow(followerId, followedId)
	if err != nil {
		return nil, err
	}
	responsePb := &followers.FollowResponse{Response: response}
	return responsePb, nil
}

func (handler *FollowersHandler) Unfollow(ctx context.Context, request *followers.UnfollowRequest) (*followers.UnfollowResponse, error) {
	//Endpoint protection
	metadata, _ := metadata.FromIncomingContext(ctx)
	sub := metadata.Get("sub")
	if sub == nil || sub[0] == "" {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized")
	}
	//Endpoint protection
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
	//Endpoint protection
	metadata, _ := metadata.FromIncomingContext(ctx)
	sub := metadata.Get("sub")
	if sub == nil || sub[0] == "" {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized")
	}
	//Endpoint protection
	followerId := request.FollowerId
	followedId := request.FollowedId
	response, err := handler.service.RemoveFollowRequest(followerId, followedId)
	if err != nil {
		return nil, err
	}
	responsePb := &followers.RemoveFollowRequestResponse{Response: response}
	return responsePb, nil
}

func (handler *FollowersHandler) Block(ctx context.Context, request *followers.Request) (*followers.Response, error) {
	//Endpoint protection
	metadata, _ := metadata.FromIncomingContext(ctx)
	sub := metadata.Get("sub")
	if sub == nil || sub[0] == "" {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized")
	}
	//Endpoint protection
	response, err := handler.service.Block(request.SubjectId, request.ObjectId)
	if err != nil {
		return nil, err
	}
	responsePb := &followers.Response{Response: response}
	return responsePb, nil
}

func (handler *FollowersHandler) Unblock(ctx context.Context, request *followers.Request) (*followers.Response, error) {
	//Endpoint protection
	metadata, _ := metadata.FromIncomingContext(ctx)
	sub := metadata.Get("sub")
	if sub == nil || sub[0] == "" {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized")
	}
	//Endpoint protection
	response, err := handler.service.Unblock(request.SubjectId, request.ObjectId)
	if err != nil {
		return nil, err
	}
	responsePb := &followers.Response{Response: response}
	return responsePb, nil
}

func (handler *FollowersHandler) GetBlockedAccounts(ctx context.Context, request *followers.Id) (*followers.IdList, error) {
	users, err := handler.service.GetBlocked(request.Id)
	if err != nil {
		return nil, err
	}
	usersPb := &followers.IdList{
		Ids: []*followers.Id{},
	}
	for _, user := range users {
		usersPb.Ids = append(usersPb.Ids, &followers.Id{Id: user.Id})
	}
	return usersPb, nil
}

func (handler *FollowersHandler) GetBlockerAccounts(ctx context.Context, request *followers.Id) (*followers.IdList, error) {
	users, err := handler.service.GetBlockers(request.Id)
	if err != nil {
		return nil, err
	}
	usersPb := &followers.IdList{
		Ids: []*followers.Id{},
	}
	for _, user := range users {
		usersPb.Ids = append(usersPb.Ids, &followers.Id{Id: user.Id})
	}
	return usersPb, nil
}
