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

func (handler *PostsHandler) PutPost(ctx context.Context, request *posts.PutPostRequest) (*posts.PutPostResponse, error) {
	post := mapNewPost(request.NewPost)
	err := handler.service.CreatePost(post)
	if err != nil {
		return nil, err
	}
	return &posts.PutPostResponse{
		Message: "Post created successfully",
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
			Id:      comment.Id.String(),
			Owner:   commentOwnerPb,
			Content: comment.Content,
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
