package api

import (
	"github.com/stojic19/XWS-TIM15/job_offers_microservice/application"
)

type JobOffersHandler struct {
	//jobOffers.UnimplementedJobOffersServiceServer
	service *application.JobOffersService
}

func NewJobOffersHandler(service *application.JobOffersService) *JobOffersHandler {
	return &JobOffersHandler{
		service: service,
	}
}
