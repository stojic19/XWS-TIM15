package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type PostsStore interface {
	GetAll() ([]*Post, error)
	GetFromUser(string) ([]*Post, error)
	Get(primitive.ObjectID) (*Post, error)
	Create(*Post) error
}
