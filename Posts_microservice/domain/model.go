package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	Id string
}

type Post struct {
	Id         primitive.ObjectID
	Title      string
	Content    string
	CreateTime time.Time
	Owner      User
	Comments   []Comment
	Likes      []User
	Dislikes   []User
}

type Comment struct {
	Id       primitive.ObjectID
	Owner    User
	Content  string
	Comments []Comment
}
