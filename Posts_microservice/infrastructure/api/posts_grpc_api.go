package api

import (
	"context"
	"github.com/stojic19/XWS-TIM15/common/proto/posts"
	"github.com/stojic19/XWS-TIM15/posts_microservice/application"
	"github.com/stojic19/XWS-TIM15/posts_microservice/domain"
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
	products, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	response := &posts.GetAllResponse{
		Posts: []*posts.Post{},
	}
	for _, product := range products {
		current := mapProduct(product)
		response.Posts = append(response.Post, current)
	}
	return response, nil
}

func mapProduct(post *domain.Post) *posts.Post {
	postPb := &pb.Product{
		Id:            post.Id.Hex(),
		Name:          post.Name,
		ClothingBrand: post.ClothingBrand,
	}
	for _, color := range post.Colors {
		postPb.Colors = append(postPb.Colors, &pb.Color{
			Code: color.Code,
			Name: color.Name,
		})
	}
	return postPb
}
