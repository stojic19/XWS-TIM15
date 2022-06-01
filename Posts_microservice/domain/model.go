package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	Id string
}

type Post struct {
	Id         primitive.ObjectID `bson:"_id"`
	Title      string             `bson:"title"`
	Content    Content            `bson:"content"`
	CreateTime time.Time          `bson:"createTime"`
	Owner      User               `bson:"owner"`
	Comments   []Comment          `bson:"comments"`
	Likes      []User             `bson:"likes"`
	Dislikes   []User             `bson:"dislikes"`
}

type Content struct {
	Text   string   `bson:"text"`
	Links  []string `bson:"links"`
	Images []string `bson:"images"`
}

type Comment struct {
	Owner      User      `bson:"owner"`
	Content    string    `bson:"content"`
	CreateTime time.Time `bson:"createTime"`
}
