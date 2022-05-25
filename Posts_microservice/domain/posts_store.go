package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostsStore interface {
	GetAll() ([]*Post, error)
	GetFromUser(string) ([]*Post, error)
	GetFromUsers([]string) ([]*Post, error)
	Get(primitive.ObjectID) (*Post, error)
	Create(*Post) error
	LikePost(primitive.ObjectID, *User) error
	DislikePost(primitive.ObjectID, *User) error
	RemoveLike(primitive.ObjectID, *User) error
	RemoveDislike(primitive.ObjectID, *User) error
	CreateComment(primitive.ObjectID, *Comment) error
}
