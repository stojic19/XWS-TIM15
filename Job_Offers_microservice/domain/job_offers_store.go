package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type JobOffersStore interface {
	GetAll() ([]*JobOffer, error)
	Get(id primitive.ObjectID) (*JobOffer, error)
	Create(*JobOffer) error
	Update(*JobOffer) error
}
