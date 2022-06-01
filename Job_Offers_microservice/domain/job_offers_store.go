package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JobOffersStore interface {
	GetAll() ([]*JobOffer, error)
	Get(id primitive.ObjectID) (*JobOffer, error)
	Create(*JobOffer) error
	Update(*JobOffer) error
	Follow(id primitive.ObjectID, user *User) error
	Unfollow(id primitive.ObjectID, user *User) error
}
