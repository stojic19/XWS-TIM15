package api

import (
	"context"
	"github.com/stojic19/XWS-TIM15/common/proto/job_offers"
	"github.com/stojic19/XWS-TIM15/job_offers_microservice/application"
	"github.com/stojic19/XWS-TIM15/job_offers_microservice/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type JobOffersHandler struct {
	job_offers.UnimplementedJobOffersServiceServer
	service *application.JobOffersService
}

func NewJobOffersHandler(service *application.JobOffersService) *JobOffersHandler {
	return &JobOffersHandler{
		service: service,
	}
}

func (handler *JobOffersHandler) GetAll(ctx context.Context, request *job_offers.GetAllRequest) (*job_offers.GetAllResponse, error) {
	metadata, _ := metadata.FromIncomingContext(ctx)
	sub := metadata.Get("sub")
	if sub == nil || sub[0] == "" {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized")
	}
	offers, err := handler.service.GetAll()
	if err != nil {
		return nil, err
	}
	offersPb := &job_offers.GetAllResponse{
		JobOffers: []*job_offers.JobOffer{},
	}
	for _, jobOffer := range offers {
		current := mapJobOffer(jobOffer)
		offersPb.JobOffers = append(offersPb.JobOffers, current)
	}
	return offersPb, nil
}

func (handler *JobOffersHandler) Get(ctx context.Context, request *job_offers.JobOfferId) (*job_offers.JobOffer, error) {
	id, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}
	offer, err := handler.service.Get(id)
	if err != nil {
		return nil, err
	}
	offerPb := mapJobOffer(offer)
	return offerPb, nil
}

func (handler *JobOffersHandler) GetSubscribed(ctx context.Context, request *job_offers.GetSubscribedRequest) (*job_offers.GetSubscribedResponse, error) {
	userId := request.Id
	offers, err := handler.service.GetSubscribed(userId)
	if err != nil {
		return nil, err
	}
	offersPb := &job_offers.GetSubscribedResponse{
		JobOffers: []*job_offers.JobOffer{},
	}
	for _, jobOffer := range offers {
		current := mapJobOffer(jobOffer)
		offersPb.JobOffers = append(offersPb.JobOffers, current)
	}
	return offersPb, nil
}

func (handler *JobOffersHandler) Create(ctx context.Context, request *job_offers.NewJobOffer) (*job_offers.Response, error) {

	metadata, _ := metadata.FromIncomingContext(ctx)
	sub := metadata.Get("sub")
	apiKey := metadata.Get("apiKey")
	if (sub == nil || sub[0] == "") && (apiKey == nil || apiKey[0] != GetAgentAppApiKey()) {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized")
	}

	jobOffer := mapNewJobOffer(request)
	err := handler.service.Create(jobOffer)
	if err != nil {
		return &job_offers.Response{
			Message: "Oops, something went wrong. Try again!",
			Code:    500,
		}, err
	}
	return &job_offers.Response{
		Message: "Job offer created!",
		Code:    200,
	}, nil
}

func (handler *JobOffersHandler) Update(ctx context.Context, request *job_offers.UpdateJobOffer) (*job_offers.Response, error) {
	jobOffer := mapJobOfferUpdate(request)
	err := handler.service.Update(jobOffer)
	if err != nil {
		return &job_offers.Response{
			Message: "Oops, something went wrong. Try again!",
			Code:    500,
		}, err
	}
	return &job_offers.Response{
		Message: "Job offer updated!",
		Code:    200,
	}, nil
}

func (handler *JobOffersHandler) FollowJobOffer(ctx context.Context, request *job_offers.SubscribeRequest) (*job_offers.Response, error) {
	jobOfferId, _ := primitive.ObjectIDFromHex(request.JobOfferId)
	user := &domain.User{Id: request.Id}
	err := handler.service.Follow(jobOfferId, user)
	if err != nil {
		return &job_offers.Response{
			Message: "Oops, something went wrong. Try again!",
			Code:    500,
		}, err
	}
	return &job_offers.Response{
		Message: "Job offer followed!",
		Code:    200,
	}, nil
}

func (handler *JobOffersHandler) UnfollowJobOffer(ctx context.Context, request *job_offers.UnsubscribeRequest) (*job_offers.Response, error) {
	jobOfferId, _ := primitive.ObjectIDFromHex(request.JobOfferId)
	user := &domain.User{Id: request.Id}
	err := handler.service.Unfollow(jobOfferId, user)
	if err != nil {
		return &job_offers.Response{
			Message: "Oops, something went wrong. Try again!",
			Code:    500,
		}, err
	}
	return &job_offers.Response{
		Message: "Job offer unfollowed!",
		Code:    200,
	}, nil
}

func mapJobOffer(jobOffer *domain.JobOffer) *job_offers.JobOffer {
	jobOfferPb := &job_offers.JobOffer{
		Id:           jobOffer.Id.Hex(),
		Position:     jobOffer.Position,
		Description:  jobOffer.Description,
		Requirements: jobOffer.Requirements,
		IsActive:     jobOffer.IsActive,
	}
	for _, follower := range jobOffer.Subscribers {
		followerPb := &job_offers.User{
			Id: follower.Id,
		}
		jobOfferPb.Subscribers = append(jobOfferPb.Subscribers, followerPb)
	}
	return jobOfferPb
}

func mapNewJobOffer(jobOffer *job_offers.NewJobOffer) *domain.JobOffer {
	domainJobOffer := &domain.JobOffer{
		Position:     jobOffer.Position,
		Description:  jobOffer.Description,
		Requirements: jobOffer.Requirements,
		Subscribers:  []domain.User{},
	}
	return domainJobOffer
}

func mapJobOfferUpdate(jobOffer *job_offers.UpdateJobOffer) *domain.JobOffer {
	id, _ := primitive.ObjectIDFromHex(jobOffer.Id)
	jobOfferPb := &domain.JobOffer{
		Id:           id,
		Position:     jobOffer.Position,
		Description:  jobOffer.Description,
		Requirements: jobOffer.Requirements,
		IsActive:     jobOffer.IsActive,
	}
	return jobOfferPb
}
