package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JobOffersStore interface {
	GetAll() ([]*JobOffer, error)
	Get(id primitive.ObjectID) (*JobOffer, error)
	GetSubscribed(string) ([]*JobOffer, error)
	Create(*JobOffer) error
	Update(*JobOffer) error
	Subscribe(id primitive.ObjectID, user *User) error
	Unsubscribe(id primitive.ObjectID, user *User) error
}
