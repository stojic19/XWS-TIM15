package application

import (
	"github.com/stojic19/XWS-TIM15/job_offers_microservice/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JobOffersService struct {
	store domain.JobOffersStore
}

func NewJobOffersService(store domain.JobOffersStore) *JobOffersService {
	return &JobOffersService{
		store: store,
	}
}

func (service *JobOffersService) GetAll() ([]*domain.JobOffer, error) {
	return service.store.GetAll()
}

func (service *JobOffersService) Get(id primitive.ObjectID) (*domain.JobOffer, error) {
	return service.store.Get(id)
}

func (service *JobOffersService) Create(offer *domain.JobOffer) error {
	return service.store.Create(offer)
}
