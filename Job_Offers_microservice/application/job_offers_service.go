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

func (service *JobOffersService) GetSubscribed(userId string) ([]*domain.JobOffer, error) {
	return service.store.GetSubscribed(userId)
}

func (service *JobOffersService) GetRecommended(userId string, skills []string) ([]*domain.JobOffer, error) {
	return service.store.GetRecommended(userId, skills)
}

func (service *JobOffersService) Create(offer *domain.JobOffer) error {
	offer.IsActive = true
	return service.store.Create(offer)
}

func (service *JobOffersService) Update(offer *domain.JobOffer) error {
	return service.store.Update(offer)
}

func (service *JobOffersService) Subscribe(jobOfferId primitive.ObjectID, user *domain.User) error {
	return service.store.Subscribe(jobOfferId, user)
}

func (service *JobOffersService) Unsubscribe(jobOfferId primitive.ObjectID, user *domain.User) error {
	return service.store.Unsubscribe(jobOfferId, user)
}
