package api

import (
	"context"
	"fmt"
	"github.com/stojic19/XWS-TIM15/Followers_microservice/application"
	"github.com/stojic19/XWS-TIM15/Followers_microservice/infrastructure/services"
	"github.com/stojic19/XWS-TIM15/common/proto/followers"
	"github.com/stojic19/XWS-TIM15/common/proto/users"
	"github.com/stojic19/XWS-TIM15/common/tracer"
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
	span := tracer.StartSpanFromContextMetadata(ctx, "GetFollows")
	defer span.Finish()

	id := request.Id
	span1 := tracer.StartSpanFromContext(tracer.ContextWithSpan(ctx, span), "Neo4jRead")
	response, err := handler.service.GetFollows(id)
	span1.Finish()
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
	span := tracer.StartSpanFromContextMetadata(ctx, "GetFollowers")
	defer span.Finish()

	id := request.Id
	span1 := tracer.StartSpanFromContext(tracer.ContextWithSpan(ctx, span), "Neo4jRead")
	response, err := handler.service.GetFollowers(id)
	span1.Finish()
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
	span := tracer.StartSpanFromContextMetadata(ctx, "GetFollowRequests")
	defer span.Finish()

	md, _ := metadata.FromIncomingContext(ctx)
	fmt.Println(md)
	id := request.Id
	span1 := tracer.StartSpanFromContext(tracer.ContextWithSpan(ctx, span), "Neo4jRead")
	response, err := handler.service.GetFollowRequests(id)
	span1.Finish()
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
	span := tracer.StartSpanFromContextMetadata(ctx, "GetFollowerRequests")
	defer span.Finish()

	id := request.Id
	span1 := tracer.StartSpanFromContext(tracer.ContextWithSpan(ctx, span), "Neo4jRead")
	response, err := handler.service.GetFollowerRequests(id)
	span1.Finish()
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
	span := tracer.StartSpanFromContextMetadata(ctx, "GetRelationship")
	defer span.Finish()

	followerId := request.FollowerId
	followedId := request.FollowedId
	span1 := tracer.StartSpanFromContext(tracer.ContextWithSpan(ctx, span), "Neo4jRead")
	response, err := handler.service.GetRelationship(followerId, followedId)
	span1.Finish()
	if err != nil {
		return nil, err
	}
	responsePb := &followers.GetRelationshipResponse{}
	responsePb.Relationship = response
	return responsePb, nil
}

func (handler *FollowersHandler) ConfirmFollow(ctx context.Context, request *followers.ConfirmFollowRequest) (*followers.ConfirmFollowResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "ConfirmFollow")
	defer span.Finish()

	//Endpoint protection
	metadata, _ := metadata.FromIncomingContext(ctx)
	sub := metadata.Get("sub")
	if sub == nil || sub[0] == "" {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized")
	}
	//Endpoint protection
	followerId := request.FollowerId
	followedId := request.FollowedId
	span1 := tracer.StartSpanFromContext(tracer.ContextWithSpan(ctx, span), "Neo4jWrite")
	response, err := handler.service.ConfirmFollow(followerId, followedId)
	span1.Finish()
	if err != nil {
		return nil, err
	}
	responsePb := &followers.ConfirmFollowResponse{Response: response}
	return responsePb, nil
}

func (handler *FollowersHandler) Follow(ctx context.Context, request *followers.FollowRequest) (*followers.FollowResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "Follow")
	defer span.Finish()

	//Endpoint protection
	metadata, _ := metadata.FromIncomingContext(ctx)
	sub := metadata.Get("sub")
	if sub == nil || sub[0] == "" {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized")
	}
	//Endpoint protection
	followerId := request.FollowerId
	followedId := request.FollowedId

	ctx = tracer.ContextWithSpan(ctx, span)
	userClient := services.NewUsersClient(handler.usersClientAddress)
	userResponse, err := userClient.GetUser(ctx, &users.GetUserRequest{Id: followedId})
	if err != nil {
		return nil, err
	}
	if userResponse.User.IsPrivate {
		span1 := tracer.StartSpanFromContext(tracer.ContextWithSpan(ctx, span), "Neo4jWriteFollowRequest")
		response, err := handler.service.FollowRequest(followerId, followedId)
		span1.Finish()
		if err != nil {
			return nil, err
		}
		responsePb := &followers.FollowResponse{Response: response}
		return responsePb, nil
	}
	span1 := tracer.StartSpanFromContext(tracer.ContextWithSpan(ctx, span), "Neo4jWriteFollow")
	response, err := handler.service.Follow(followerId, followedId)
	span1.Finish()
	if err != nil {
		return nil, err
	}
	responsePb := &followers.FollowResponse{Response: response}
	return responsePb, nil
}

func (handler *FollowersHandler) Unfollow(ctx context.Context, request *followers.UnfollowRequest) (*followers.UnfollowResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "Unfollow")
	defer span.Finish()

	//Endpoint protection
	metadata, _ := metadata.FromIncomingContext(ctx)
	sub := metadata.Get("sub")
	if sub == nil || sub[0] == "" {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized")
	}
	//Endpoint protection
	followerId := request.FollowerId
	followedId := request.FollowedId
	span1 := tracer.StartSpanFromContext(tracer.ContextWithSpan(ctx, span), "Neo4jWriteUnfollow")
	response, err := handler.service.Unfollow(followerId, followedId)
	span1.Finish()
	if err != nil {
		return nil, err
	}
	responsePb := &followers.UnfollowResponse{Response: response}
	return responsePb, nil
}

func (handler *FollowersHandler) RemoveFollowRequest(ctx context.Context, request *followers.RemoveFollowRequestRequest) (*followers.RemoveFollowRequestResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "RemoveFollowRequest")
	defer span.Finish()

	//Endpoint protection
	metadata, _ := metadata.FromIncomingContext(ctx)
	sub := metadata.Get("sub")
	if sub == nil || sub[0] == "" {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized")
	}
	//Endpoint protection
	followerId := request.FollowerId
	followedId := request.FollowedId
	span1 := tracer.StartSpanFromContext(tracer.ContextWithSpan(ctx, span), "Neo4jWriteRemoveFollowRequest")
	response, err := handler.service.RemoveFollowRequest(followerId, followedId)
	span1.Finish()
	if err != nil {
		return nil, err
	}
	responsePb := &followers.RemoveFollowRequestResponse{Response: response}
	return responsePb, nil
}

func (handler *FollowersHandler) Block(ctx context.Context, request *followers.Request) (*followers.Response, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "Block")
	defer span.Finish()

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
	span := tracer.StartSpanFromContextMetadata(ctx, "Unblock")
	defer span.Finish()

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
	span := tracer.StartSpanFromContextMetadata(ctx, "GetBlockedAccounts")
	defer span.Finish()

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
	span := tracer.StartSpanFromContextMetadata(ctx, "GetBlockerAccounts")
	defer span.Finish()

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

func (handler *FollowersHandler) GetRecommendedUsers(ctx context.Context, request *followers.Id) (*followers.IdList, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetRecommendedUsers")
	defer span.Finish()

	users, err := handler.service.GetRecommendedUsers(request.Id)
	if err != nil {
		return nil, err
	}
	retVal := &followers.IdList{
		Ids: []*followers.Id{},
	}
	for _, user := range users {
		retVal.Ids = append(retVal.Ids, &followers.Id{Id: user.Id})
	}
	return retVal, nil
}
