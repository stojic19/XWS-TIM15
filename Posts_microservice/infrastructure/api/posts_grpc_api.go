package api

import (
	"context"
	"github.com/stojic19/XWS-TIM15/common/proto/posts"
	"github.com/stojic19/XWS-TIM15/posts_microservice/application"
	"github.com/stojic19/XWS-TIM15/posts_microservice/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type PostsHandler struct {
	posts.UnimplementedPostsServiceServer
	service *application.PostsService
}

func NewPostsHandler(service *application.PostsService) *PostsHandler {
	return &PostsHandler{
		service: service,
	}
}

func (handler *PostsHandler) GetAll(ctx context.Context, request *posts.GetAllRequest) (*posts.GetAllResponse, error) {
	returnPosts, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &posts.GetAllResponse{
		Posts: []*posts.Post{},
	}
	for _, post := range returnPosts {
		current := mapPost(post)
		response.Posts = append(response.Posts, current)
	}
	return response, nil
}

func (handler *PostsHandler) Get(ctx context.Context, request *posts.GetRequest) (*posts.GetResponse, error) {
	id := request.Id
	postId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	post, err := handler.service.Get(postId)
	if err != nil {
		return nil, err
	}
	postPb := mapPost(post)
	response := &posts.GetResponse{
		Post: postPb,
	}
	return response, nil
}

func (handler *PostsHandler) GetFromUser(ctx context.Context, request *posts.GetFromUserRequest) (*posts.GetFromUserResponse, error) {
	id := request.Id
	returnPosts, err := handler.service.GetFromUser(id)
	if err != nil {
		return nil, err
	}
	response := &posts.GetFromUserResponse{
		Posts: []*posts.Post{},
	}
	for _, post := range returnPosts {
		current := mapPost(post)
		response.Posts = append(response.Posts, current)
	}
	return response, nil
}

func (handler *PostsHandler) GetFromFollowed(ctx context.Context, request *posts.GetFollowedRequest) (*posts.GetFollowedResponse, error) {
	id := request.Id
	// treba promeniti na one koje user prati
	returnPosts, err := handler.service.GetFromUser(id)
	if err != nil {
		return nil, err
	}
	response := &posts.GetFollowedResponse{
		Posts: []*posts.Post{},
	}
	for _, post := range returnPosts {
		current := mapPost(post)
		response.Posts = append(response.Posts, current)
	}
	return response, nil
}

func (handler *PostsHandler) CreatePost(ctx context.Context, request *posts.CreatePostRequest) (*posts.CreatePostResponse, error) {
	post := mapNewPost(request.NewPost)
	err := handler.service.CreatePost(post)
	if err != nil {
		return nil, err
	}
	return &posts.CreatePostResponse{
		Message: "Post created successfully",
	}, nil
}

func (handler *PostsHandler) LikePost(ctx context.Context, request *posts.LikePostRequest) (*posts.LikePostResponse, error) {
	postId, err := primitive.ObjectIDFromHex(request.PostId)
	if err != nil {
		return nil, err
	}
	userId := mapNewUser(request.UserId)
	err = handler.service.LikePost(postId, userId)
	if err != nil {
		return nil, err
	}
	return &posts.LikePostResponse{
		Message: "Like successful",
	}, nil
}

func (handler *PostsHandler) RemoveLike(ctx context.Context, request *posts.RemoveLikeRequest) (*posts.RemoveLikeResponse, error) {
	postId, err := primitive.ObjectIDFromHex(request.PostId)
	if err != nil {
		return nil, err
	}
	userId := mapNewUser(request.UserId)
	err = handler.service.RemoveLike(postId, userId)
	if err != nil {
		return nil, err
	}
	return &posts.RemoveLikeResponse{
		Message: "Like removed successfully",
	}, nil
}

func (handler *PostsHandler) DislikePost(ctx context.Context, request *posts.DislikePostRequest) (*posts.DislikePostResponse, error) {
	postId, err := primitive.ObjectIDFromHex(request.PostId)
	if err != nil {
		return nil, err
	}
	userId := mapNewUser(request.UserId)
	err = handler.service.DislikePost(postId, userId)
	if err != nil {
		return nil, err
	}
	return &posts.DislikePostResponse{
		Message: "Dislike successful",
	}, nil
}

func (handler *PostsHandler) RemoveDislike(ctx context.Context, request *posts.RemoveDislikeRequest) (*posts.RemoveDislikeResponse, error) {
	postId, err := primitive.ObjectIDFromHex(request.PostId)
	if err != nil {
		return nil, err
	}
	userId := mapNewUser(request.UserId)
	err = handler.service.RemoveLike(postId, userId)
	if err != nil {
		return nil, err
	}
	return &posts.RemoveDislikeResponse{
		Message: "Dislike removed successfully",
	}, nil
}

func (handler *PostsHandler) CommentPost(ctx context.Context, request *posts.CommentPostRequest) (*posts.CommentPostResponse, error) {
	postId, err := primitive.ObjectIDFromHex(request.PostId)
	if err != nil {
		return nil, err
	}
	comment := mapNewComment(request.UserId, request.Content)
	err = handler.service.CreateComment(postId, comment)
	if err != nil {
		return nil, err
	}
	return &posts.CommentPostResponse{
		Message: "Comment created successfully",
	}, nil
}

func mapPost(post *domain.Post) *posts.Post {
	ownerPb := &posts.User{
		Id: post.Owner.Id,
	}
	postPb := &posts.Post{
		Id:         post.Id.Hex(),
		Title:      post.Title,
		Content:    post.Content,
		CreateTime: timestamppb.New(post.CreateTime),
		Owner:      ownerPb,
	}
	for _, comment := range post.Comments {
		commentOwnerPb := &posts.User{
			Id: comment.Owner.Id,
		}
		currentCommentPb := &posts.Comment{
			//Id:      comment.Id.String(),
			Owner:      commentOwnerPb,
			Content:    comment.Content,
			CreateTime: timestamppb.New(comment.CreateTime),
		}
		postPb.Comments = append(postPb.Comments, currentCommentPb)
	}
	for _, userLike := range post.Likes {
		likePb := &posts.User{
			Id: userLike.Id,
		}
		postPb.Likes = append(postPb.Likes, likePb)
	}
	for _, userDislike := range post.Dislikes {
		dislikePb := &posts.User{
			Id: userDislike.Id,
		}
		postPb.Dislikes = append(postPb.Dislikes, dislikePb)
	}

	return postPb
}

func mapNewPost(post *posts.NewPost) *domain.Post {
	newPostOwner := domain.User{
		Id: post.Owner.Id,
	}
	newPost := &domain.Post{
		Title:      post.Title,
		Content:    post.Content,
		Owner:      newPostOwner,
		CreateTime: time.Now(),
		Comments:   []domain.Comment{},
		Likes:      []domain.User{},
		Dislikes:   []domain.User{},
	}
	return newPost
}

func mapNewUser(id string) *domain.User {
	user := &domain.User{
		Id: id,
	}
	return user
}

func mapNewComment(userId string, content string) *domain.Comment {
	comment := &domain.Comment{
		Owner:      *mapNewUser(userId),
		Content:    content,
		CreateTime: time.Now(),
	}
	return comment
}