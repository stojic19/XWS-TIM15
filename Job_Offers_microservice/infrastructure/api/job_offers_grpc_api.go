package api

import (
	"context"
	otgo "github.com/opentracing/opentracing-go"
	"github.com/stojic19/XWS-TIM15/common/proto/job_offers"
	"github.com/stojic19/XWS-TIM15/common/proto/users"
	"github.com/stojic19/XWS-TIM15/common/tracer"
	"github.com/stojic19/XWS-TIM15/job_offers_microservice/application"
	"github.com/stojic19/XWS-TIM15/job_offers_microservice/domain"
	"github.com/stojic19/XWS-TIM15/job_offers_microservice/infrastructure/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type JobOffersHandler struct {
	job_offers.UnimplementedJobOffersServiceServer
	service            *application.JobOffersService
	usersClientAddress string
}

func NewJobOffersHandler(service *application.JobOffersService, usersEndpoint string) *JobOffersHandler {
	return &JobOffersHandler{
		service:            service,
		usersClientAddress: usersEndpoint,
	}
}

func (handler *JobOffersHandler) GetAll(ctx context.Context, request *job_offers.GetAllRequest) (*job_offers.GetAllResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAll")
	defer span.Finish()

	span1 := tracer.StartSpanFromContext(tracer.ContextWithSpan(ctx, span), "MongoReadGetAll")
	offers, err := handler.service.GetAll()
	span1.Finish()
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
	span := tracer.StartSpanFromContextMetadata(ctx, "Get")
	defer span.Finish()

	id, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, err
	}
	span1 := tracer.StartSpanFromContext(tracer.ContextWithSpan(ctx, span), "MongoGet")
	offer, err := handler.service.Get(id)
	span1.Finish()
	if err != nil {
		return nil, err
	}
	offerPb := mapJobOffer(offer)
	return offerPb, nil
}

func (handler *JobOffersHandler) GetSubscribed(ctx context.Context, request *job_offers.GetSubscribedRequest) (*job_offers.GetSubscribedResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetSubscribed")
	defer span.Finish()

	userId := request.Id
	span1 := tracer.StartSpanFromContext(tracer.ContextWithSpan(ctx, span), "MongoGet")
	offers, err := handler.service.GetSubscribed(userId)
	span1.Finish()
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
	span := tracer.StartSpanFromContextMetadata(ctx, "Create")
	defer span.Finish()

	metadata, _ := metadata.FromIncomingContext(ctx)
	sub := metadata.Get("sub")
	apiKey := metadata.Get("apiKey")

	ctx1 := tracer.InjectToMetadata(ctx, otgo.GlobalTracer(), span)
	userClient := services.NewUsersClient(handler.usersClientAddress)
	response, err := userClient.ValidateApiKey(ctx1, &users.ApiKey{ApiKey: apiKey[0]})
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized")
	}
	if (sub == nil || sub[0] == "") && (response.IsValid == false) {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized")
	}

	jobOffer := mapNewJobOffer(request)
	span1 := tracer.StartSpanFromContext(tracer.ContextWithSpan(ctx, span), "MongoCreate")
	err = handler.service.Create(jobOffer)
	span1.Finish()
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
	span := tracer.StartSpanFromContextMetadata(ctx, "Update")
	defer span.Finish()

	//Endpoint protection
	metadata, _ := metadata.FromIncomingContext(ctx)
	sub := metadata.Get("sub")
	if sub == nil || sub[0] == "" {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized")
	}
	//Endpoint protection
	jobOffer := mapJobOfferUpdate(request)
	span1 := tracer.StartSpanFromContext(tracer.ContextWithSpan(ctx, span), "MongoUpdate")
	err := handler.service.Update(jobOffer)
	span1.Finish()
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

func (handler *JobOffersHandler) SubscribeJobOffer(ctx context.Context, request *job_offers.SubscribeRequest) (*job_offers.Response, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "SubscribeJobOffer")
	defer span.Finish()

	//Endpoint protection
	metadata, _ := metadata.FromIncomingContext(ctx)
	sub := metadata.Get("sub")
	if sub == nil || sub[0] == "" {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized")
	}
	//Endpoint protection
	jobOfferId, _ := primitive.ObjectIDFromHex(request.JobOfferId)
	user := &domain.User{Id: request.Id}
	span1 := tracer.StartSpanFromContext(tracer.ContextWithSpan(ctx, span), "MongoSubscribe")
	err := handler.service.Subscribe(jobOfferId, user)
	span1.Finish()
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

func (handler *JobOffersHandler) UnsubscribeJobOffer(ctx context.Context, request *job_offers.UnsubscribeRequest) (*job_offers.Response, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "UnsubscribeJobOffer")
	defer span.Finish()

	//Endpoint protection
	metadata, _ := metadata.FromIncomingContext(ctx)
	sub := metadata.Get("sub")
	if sub == nil || sub[0] == "" {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized")
	}
	//Endpoint protection
	jobOfferId, _ := primitive.ObjectIDFromHex(request.JobOfferId)
	user := &domain.User{Id: request.Id}
	span1 := tracer.StartSpanFromContext(tracer.ContextWithSpan(ctx, span), "MongoUnsubscribe")
	err := handler.service.Unsubscribe(jobOfferId, user)
	span1.Finish()
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
