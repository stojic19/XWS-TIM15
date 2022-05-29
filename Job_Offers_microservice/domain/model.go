package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type JobOffer struct {
	Id           primitive.ObjectID `bson:"_id"`
	Position     string             `bson:"position"`
	Description  string             `bson:"description"`
	Requirements string             `bson:"requirements"`
	IsActive     bool               `bson:"isActive"`
}
