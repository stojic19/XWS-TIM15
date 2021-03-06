package api

import (
	"context"
	"encoding/json"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	otgo "github.com/opentracing/opentracing-go"
	"github.com/stojic19/XWS-TIM15/api_gateway/domain"
	"github.com/stojic19/XWS-TIM15/api_gateway/infrastructure/services"
	"github.com/stojic19/XWS-TIM15/common/proto/followers"
	"github.com/stojic19/XWS-TIM15/common/proto/users"
	"github.com/stojic19/XWS-TIM15/common/tracer"
	"google.golang.org/grpc/metadata"
	"net/http"
)

type FollowersHandler struct {
	followersClientAddress string
	usersClientAddress     string
}

func NewFollowersHandler(followersClientAddress, usersClientAddress string) Handler {
	return &FollowersHandler{
		followersClientAddress: followersClientAddress,
		usersClientAddress:     usersClientAddress,
	}
}

func (handler *FollowersHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/followers/followers/{userId}/details", handler.GetFollowersDetails)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("GET", "/followers/follows/{userId}/details", handler.GetFollowsDetails)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("GET", "/followers/followRequests/{userId}/details", handler.GetFollowerRequestsDetails)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("GET", "/followers/followerRequests/{userId}/details", handler.GetFollowRequestsDetails)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("GET", "/followers/recommended/{userId}/details", handler.GetRecommended)
	if err != nil {
		panic(err)
	}
}

func (handler *FollowersHandler) GetFollowersDetails(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	span := tracer.StartSpanFromRequest("GetFollowersDetails", otgo.GlobalTracer(), r)
	defer span.Finish()
	ctx := tracer.InjectToMetadata(context.TODO(), otgo.GlobalTracer(), span)

	userId, followersInfo, error := initializeFollowers(w, pathParams)
	if error {
		return
	}

	err := handler.addFollowerInfo(followersInfo, userId, ctx)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	for _, followerInfo := range followersInfo.Users {
		handler.addUserInfo(followerInfo, ctx)
		handler.addFollowedRelationship(followerInfo, userId, ctx)
	}

	finishFollowers(w, err, followersInfo)
}

func (handler *FollowersHandler) GetFollowsDetails(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	span := tracer.StartSpanFromRequest("GetFollowsDetails", otgo.GlobalTracer(), r)
	defer span.Finish()
	ctx := tracer.InjectToMetadata(context.TODO(), otgo.GlobalTracer(), span)

	userId, followersInfo, error := initializeFollowers(w, pathParams)
	if error {
		return
	}

	err := handler.addFollowInfo(followersInfo, userId, ctx)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	for _, followerInfo := range followersInfo.Users {
		handler.addUserInfo(followerInfo, ctx)
		handler.addFollowerRelationship(followerInfo, userId, ctx)
	}

	finishFollowers(w, err, followersInfo)
}

func (handler *FollowersHandler) GetFollowerRequestsDetails(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	span := tracer.StartSpanFromRequest("GetFollowerRequestsDetails", otgo.GlobalTracer(), r)
	defer span.Finish()
	ctx := tracer.InjectToMetadata(context.TODO(), otgo.GlobalTracer(), span)

	userId, followersInfo, error := initializeFollowers(w, pathParams)
	if error {
		return
	}

	err := handler.addFollowerRequestInfo(followersInfo, userId, ctx)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	for _, followerInfo := range followersInfo.Users {
		handler.addUserInfo(followerInfo, ctx)
		handler.addFollowedRelationship(followerInfo, userId, ctx)
	}

	finishFollowers(w, err, followersInfo)
}

func (handler *FollowersHandler) GetFollowRequestsDetails(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	span := tracer.StartSpanFromRequest("GetFollowRequestsDetails", otgo.GlobalTracer(), r)
	defer span.Finish()
	ctx := tracer.InjectToMetadata(context.TODO(), otgo.GlobalTracer(), span)

	userId, followersInfo, error := initializeFollowers(w, pathParams)
	if error {
		return
	}

	err := handler.addFollowRequestInfo(followersInfo, userId, ctx)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	for _, followerInfo := range followersInfo.Users {
		handler.addUserInfo(followerInfo, ctx)
		handler.addFollowerRelationship(followerInfo, userId, ctx)
	}

	finishFollowers(w, err, followersInfo)
}

func (handler *FollowersHandler) GetRecommended(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	span := tracer.StartSpanFromRequest("GetRecommended", otgo.GlobalTracer(), r)
	defer span.Finish()

	ctx := tracer.InjectToMetadata(context.TODO(), otgo.GlobalTracer(), span)
	followersClient := services.NewFollowersClient(handler.followersClientAddress)
	ids, err := followersClient.GetRecommendedUsers(ctx, &followers.Id{Id: pathParams["userId"]})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	usersClient := services.NewUsersClient(handler.usersClientAddress)
	retVal := []*domain.User{}
	for _, id := range ids.Ids {
		info, err := usersClient.GetUser(ctx, &users.GetUserRequest{Id: id.Id})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		retVal = append(retVal, &domain.User{
			Id:       id.Id,
			Username: info.User.Username,
		})
	}
	retValByte, err := json.Marshal(retVal)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(retValByte)
}

func (handler *FollowersHandler) addFollowerInfo(followersInfo *domain.UserFollowerInfoList, id string, context context.Context) error {
	followersClient := services.NewFollowersClient(handler.followersClientAddress)
	followersIds, err := followersClient.GetFollowers(context, &followers.GetFollowersRequest{Id: id})
	if err != nil {
		return err
	}
	for _, follower := range followersIds.Followers {
		followerInfo := domain.UserFollowerInfo{
			Id:        follower.Id,
			StartDate: follower.Time.AsTime(),
		}
		followersInfo.Users = append(followersInfo.Users, &followerInfo)
	}
	return nil
}

func (handler *FollowersHandler) addFollowInfo(followersInfo *domain.UserFollowerInfoList, id string, context context.Context) error {
	followersClient := services.NewFollowersClient(handler.followersClientAddress)
	followsIds, err := followersClient.GetFollows(context, &followers.GetFollowsRequest{Id: id})
	if err != nil {
		return err
	}
	for _, followed := range followsIds.Follows {
		followerInfo := domain.UserFollowerInfo{
			Id:        followed.Id,
			StartDate: followed.Time.AsTime(),
		}
		followersInfo.Users = append(followersInfo.Users, &followerInfo)
	}
	return nil
}

func (handler *FollowersHandler) addFollowerRequestInfo(followersInfo *domain.UserFollowerInfoList, id string, context context.Context) error {
	followersClient := services.NewFollowersClient(handler.followersClientAddress)
	followerRequestsIds, err := followersClient.GetFollowerRequests(context, &followers.GetFollowerRequestsRequest{Id: id})
	if err != nil {
		return err
	}
	for _, followerRequest := range followerRequestsIds.FollowerRequests {
		followerInfo := domain.UserFollowerInfo{
			Id:        followerRequest.Id,
			StartDate: followerRequest.Time.AsTime(),
		}
		followersInfo.Users = append(followersInfo.Users, &followerInfo)
	}
	return nil
}

func (handler *FollowersHandler) addFollowRequestInfo(followersInfo *domain.UserFollowerInfoList, id string, ctx context.Context) error {
	followersClient := services.NewFollowersClient(handler.followersClientAddress)
	metadata1 := metadata.MD{}
	metadata1.Set("proba", "1")
	context1 := context.TODO()
	context1 = metadata.NewOutgoingContext(context1, metadata1)
	followRequestsIds, err := followersClient.GetFollowRequests(ctx, &followers.GetFollowRequestsRequest{Id: id})
	if err != nil {
		return err
	}
	for _, followRequest := range followRequestsIds.FollowRequests {
		followerInfo := domain.UserFollowerInfo{
			Id:        followRequest.Id,
			StartDate: followRequest.Time.AsTime(),
		}
		followersInfo.Users = append(followersInfo.Users, &followerInfo)
	}
	return nil
}

func (handler *FollowersHandler) addUserInfo(followerInfo *domain.UserFollowerInfo, context context.Context) {
	usersClient := services.NewUsersClient(handler.usersClientAddress)
	userInfo, err := usersClient.GetUser(context, &users.GetUserRequest{Id: followerInfo.Id})
	if err != nil {
		return
	}
	followerInfo.Username = userInfo.User.Username
	followerInfo.Name = userInfo.User.Name
	followerInfo.Gender = userInfo.User.Gender
}

func (handler *FollowersHandler) addFollowerRelationship(followerInfo *domain.UserFollowerInfo, mainId string, context context.Context) {
	followersClient := services.NewFollowersClient(handler.followersClientAddress)
	relationship, err := followersClient.GetRelationship(context, &followers.GetRelationshipRequest{FollowedId: mainId, FollowerId: followerInfo.Id})
	if err != nil {
		return
	}
	followerInfo.ReverseRelationship = relationship.Relationship
}

func (handler *FollowersHandler) addFollowedRelationship(followedInfo *domain.UserFollowerInfo, mainId string, context context.Context) {
	followersClient := services.NewFollowersClient(handler.followersClientAddress)
	relationship, err := followersClient.GetRelationship(context, &followers.GetRelationshipRequest{FollowedId: followedInfo.Id, FollowerId: mainId})
	if err != nil {
		return
	}
	followedInfo.ReverseRelationship = relationship.Relationship
}

func initializeFollowers(w http.ResponseWriter, pathParams map[string]string) (string, *domain.UserFollowerInfoList, bool) {
	id := pathParams["userId"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return "", nil, true
	}
	followersInfo := &domain.UserFollowerInfoList{}
	followersInfo.Users = []*domain.UserFollowerInfo{}
	return id, followersInfo, false
}

func finishFollowers(w http.ResponseWriter, err error, followersInfo *domain.UserFollowerInfoList) {
	response, err := json.Marshal(followersInfo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
