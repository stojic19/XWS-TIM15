package api

import (
	"context"
	"encoding/json"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	otgo "github.com/opentracing/opentracing-go"
	"github.com/stojic19/XWS-TIM15/api_gateway/domain"
	"github.com/stojic19/XWS-TIM15/api_gateway/infrastructure/services"
	"github.com/stojic19/XWS-TIM15/common/proto/followers"
	"github.com/stojic19/XWS-TIM15/common/proto/posts"
	"github.com/stojic19/XWS-TIM15/common/proto/users"
	"github.com/stojic19/XWS-TIM15/common/tracer"
	"net/http"
	"time"
)

type PostsHandler struct {
	postsClientAddress     string
	followersClientAddress string
	usersClientAddress     string
}

func NewPostsHandler(postsClientAddress, followersClientAddress, usersClientAddress string) Handler {
	return &PostsHandler{
		postsClientAddress:     postsClientAddress,
		followersClientAddress: followersClientAddress,
		usersClientAddress:     usersClientAddress,
	}
}

func (handler *PostsHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/posts/posts/details", handler.GetAllPostsDetails)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("GET", "/posts/posts/{postId}/details", handler.GetPostDetails)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("GET", "/posts/postsFromUser/{userId}/details", handler.GetPostFromUserDetails)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("GET", "/posts/postsFollowed/{userId}/details", handler.GetPostFromFollowedDetails)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("GET", "/posts/public/details", handler.GetPostsFromPublicDetails)
	if err != nil {
		panic(err)
	}
}

func (handler *PostsHandler) GetAllPostsDetails(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	span := tracer.StartSpanFromRequest("GetAllPostsDetails", otgo.GlobalTracer(), r)
	defer span.Finish()
	ctx := tracer.InjectToMetadata(context.TODO(), otgo.GlobalTracer(), span)

	postsInfo, error := initializePosts(w, pathParams)
	if error {
		return
	}

	err := handler.addPosts(postsInfo, ctx)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	for _, postInfo := range postsInfo.Posts {
		insertSideInfo(handler, postInfo, ctx)
	}

	finishPosts(w, err, postsInfo)
}

func (handler *PostsHandler) GetPostFromUserDetails(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	span := tracer.StartSpanFromRequest("GetPostFromUserDetails", otgo.GlobalTracer(), r)
	defer span.Finish()
	ctx := tracer.InjectToMetadata(context.TODO(), otgo.GlobalTracer(), span)

	userId, postsInfo, error := initializePostsWithParam(w, pathParams)
	if error {
		return
	}

	err := handler.addPostsFromUser(postsInfo, userId, ctx)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	for _, postInfo := range postsInfo.Posts {
		insertSideInfo(handler, postInfo, ctx)
	}

	finishPosts(w, err, postsInfo)
}

func (handler *PostsHandler) GetPostFromFollowedDetails(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	span := tracer.StartSpanFromRequest("GetPostFromFollowedDetails", otgo.GlobalTracer(), r)
	defer span.Finish()
	ctx := tracer.InjectToMetadata(context.TODO(), otgo.GlobalTracer(), span)

	userId, postsInfo, error := initializePostsWithParam(w, pathParams)
	if error {
		return
	}

	err := handler.addPostsFromFollows(postsInfo, userId, ctx)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	for _, postInfo := range postsInfo.Posts {
		insertSideInfo(handler, postInfo, ctx)
	}

	finishPosts(w, err, postsInfo)
}

func (handler *PostsHandler) GetPostDetails(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	span := tracer.StartSpanFromRequest("GetPostDetails", otgo.GlobalTracer(), r)
	defer span.Finish()
	ctx := tracer.InjectToMetadata(context.TODO(), otgo.GlobalTracer(), span)

	postId, postInfo, error := initializePost(w, pathParams)
	if error {
		return
	}

	err := handler.addPost(postInfo, postId, ctx)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	handler.addUserInfo(postInfo.Owner, ctx)
	insertSideInfo(handler, postInfo, ctx)

	finishPost(w, err, postInfo)
}

func (handler *PostsHandler) GetPostsFromPublicDetails(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	span := tracer.StartSpanFromRequest("GetPostsFromPublicDetails", otgo.GlobalTracer(), r)
	defer span.Finish()
	ctx := tracer.InjectToMetadata(context.TODO(), otgo.GlobalTracer(), span)

	postsInfo, error := initializePosts(w, pathParams)
	if error {
		return
	}

	err := handler.addPublicPosts(postsInfo, ctx)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	for _, postInfo := range postsInfo.Posts {
		insertSideInfo(handler, postInfo, ctx)
	}

	finishPosts(w, err, postsInfo)
}

func (handler *PostsHandler) addPosts(postsInfo *domain.PostUsersInfoList, ctx context.Context) error {
	postsClient := services.NewPostsClient(handler.postsClientAddress)
	postsList, err := postsClient.GetAll(ctx, &posts.GetAllRequest{})
	if err != nil {
		return err
	}
	for _, post := range postsList.Posts {
		postInfo := mapPost(post)
		postsInfo.Posts = append(postsInfo.Posts, &postInfo)
	}
	return nil
}

func (handler *PostsHandler) addPublicPosts(postsInfo *domain.PostUsersInfoList, ctx context.Context) error {
	postsClient := services.NewPostsClient(handler.postsClientAddress)
	postsList, err := postsClient.GetFromPublic(ctx, &posts.GetPublicRequest{})
	if err != nil {
		return err
	}
	for _, post := range postsList.Posts {
		postInfo := mapPost(post)
		postsInfo.Posts = append(postsInfo.Posts, &postInfo)
	}
	return nil
}

func (handler *PostsHandler) addPostsFromUser(postsInfo *domain.PostUsersInfoList, userId string, ctx context.Context) error {
	postsClient := services.NewPostsClient(handler.postsClientAddress)
	postsList, err := postsClient.GetFromUser(ctx, &posts.GetFromUserRequest{Id: userId})
	if err != nil {
		return err
	}
	for _, post := range postsList.Posts {
		postInfo := mapPost(post)
		postsInfo.Posts = append(postsInfo.Posts, &postInfo)
	}
	return nil
}

func (handler *PostsHandler) addPostsFromFollows(postsInfo *domain.PostUsersInfoList, userId string, ctx context.Context) error {
	postsClient := services.NewPostsClient(handler.postsClientAddress)
	postsList, err := postsClient.GetFromFollowed(ctx, &posts.GetFollowedRequest{Id: userId})
	if err != nil {
		return err
	}
	for _, post := range postsList.Posts {
		postInfo := mapPost(post)
		postsInfo.Posts = append(postsInfo.Posts, &postInfo)
	}
	return nil
}

func (handler *PostsHandler) addPost(postInfo *domain.PostUsersInfo, postId string, ctx context.Context) error {
	postsClient := services.NewPostsClient(handler.postsClientAddress)
	postResponse, err := postsClient.Get(ctx, &posts.GetRequest{Id: postId})
	if err != nil {
		return err
	}
	*postInfo = mapPost(postResponse.Post)
	return nil
}

func (handler *PostsHandler) addUserInfo(userInfo *domain.UserPostInfo, ctx context.Context) error {
	userClient := services.NewUsersClient(handler.usersClientAddress)
	userResponse, err := userClient.GetUser(ctx, &users.GetUserRequest{Id: userInfo.Id})
	if err != nil {
		return err
	}
	userInfo.Name = userResponse.User.Name
	userInfo.Username = userResponse.User.Username
	userInfo.Gender = userResponse.User.Gender
	userInfo.DateOfBirth, _ = time.Parse("MM/DD/YYYY", userResponse.User.DateOfBirth)
	return nil
}

func (handler *PostsHandler) addRelationships(userInfo *domain.UserPostInfo, ownerId string, ctx context.Context) error {
	followersClient := services.NewFollowersClient(handler.followersClientAddress)
	outgoingRelationship, err := followersClient.GetRelationship(ctx, &followers.GetRelationshipRequest{FollowedId: userInfo.Id, FollowerId: ownerId})
	if err != nil {
		return err
	}
	ingoingRelationship, err := followersClient.GetRelationship(ctx, &followers.GetRelationshipRequest{FollowedId: ownerId, FollowerId: userInfo.Id})
	if err != nil {
		return err
	}
	userInfo.OutgoingRelationship = outgoingRelationship.Relationship
	userInfo.IngoingRelationship = ingoingRelationship.Relationship
	return nil
}

func insertSideInfo(handler *PostsHandler, postInfo *domain.PostUsersInfo, ctx context.Context) {
	handler.addUserInfo(postInfo.Owner, ctx)
	for _, comment := range postInfo.Comments {
		handler.addUserInfo(comment.Owner, ctx)
		handler.addRelationships(comment.Owner, postInfo.Owner.Id, ctx)
	}
	for _, userInfo := range postInfo.Likes {
		handler.addUserInfo(userInfo, ctx)
		handler.addRelationships(userInfo, postInfo.Owner.Id, ctx)
	}
	for _, userInfo := range postInfo.Dislikes {
		handler.addUserInfo(userInfo, ctx)
		handler.addRelationships(userInfo, postInfo.Owner.Id, ctx)
	}
}

func mapPost(post *posts.Post) domain.PostUsersInfo {
	postInfo := domain.PostUsersInfo{
		Id:    post.Id,
		Title: post.Title,
		Content: domain.Content{
			Text:   post.Content.Text,
			Links:  post.Content.Links,
			Images: post.Content.Images,
		},
		CreateTime: post.CreateTime.AsTime(),
		Owner: &domain.UserPostInfo{
			Id: post.Owner.Id,
		},
	}
	for _, comment := range post.Comments {
		commentInfo := domain.CommentPostInfo{
			Content:    comment.Content,
			CreateTime: comment.CreateTime.AsTime(),
			Owner: &domain.UserPostInfo{
				Id: comment.Owner.Id,
			},
		}
		postInfo.Comments = append(postInfo.Comments, &commentInfo)
	}
	for _, like := range post.Likes {
		likeInfo := domain.UserPostInfo{
			Id: like.Id,
		}
		postInfo.Likes = append(postInfo.Likes, &likeInfo)
	}
	for _, dislike := range post.Dislikes {
		dislikeInfo := domain.UserPostInfo{
			Id: dislike.Id,
		}
		postInfo.Dislikes = append(postInfo.Dislikes, &dislikeInfo)
	}
	return postInfo
}

func initializePosts(w http.ResponseWriter, pathParams map[string]string) (*domain.PostUsersInfoList, bool) {
	postsInfo := &domain.PostUsersInfoList{}
	postsInfo.Posts = []*domain.PostUsersInfo{}
	return postsInfo, false
}

func initializePostsWithParam(w http.ResponseWriter, pathParams map[string]string) (string, *domain.PostUsersInfoList, bool) {
	id := pathParams["userId"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return "", nil, true
	}

	postsInfo := &domain.PostUsersInfoList{}
	postsInfo.Posts = []*domain.PostUsersInfo{}
	return id, postsInfo, false
}

func initializePost(w http.ResponseWriter, pathParams map[string]string) (string, *domain.PostUsersInfo, bool) {
	id := pathParams["postId"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return "", nil, true
	}
	postInfo := &domain.PostUsersInfo{}
	return id, postInfo, false
}

func finishPosts(w http.ResponseWriter, err error, postsInfo *domain.PostUsersInfoList) {
	response, err := json.Marshal(postsInfo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func finishPost(w http.ResponseWriter, err error, postInfo *domain.PostUsersInfo) {
	response, err := json.Marshal(postInfo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
