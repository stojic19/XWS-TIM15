package persistence

import (
	"context"
	"github.com/stojic19/XWS-TIM15/posts_microservice/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "posts"
	COLLECTION = "posts"
)

type PostsMongoStore struct {
	posts *mongo.Collection
}

func NewPostsStore(client *mongo.Client) domain.PostsStore {
	posts := client.Database(DATABASE).Collection(COLLECTION)
	return &PostsMongoStore{
		posts: posts,
	}
}

func (store *PostsMongoStore) GetAll() ([]*domain.Post, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *PostsMongoStore) Create(post *domain.Post) error {
	post.Id = primitive.NewObjectID()
	result, err := store.posts.InsertOne(context.TODO(), post)
	if err != nil {
		return err
	}
	post.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *PostsMongoStore) filter(filter interface{}) ([]*domain.Post, error) {
	cursor, err := store.posts.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *PostsMongoStore) filterOne(filter interface{}) (product *domain.Post, err error) {
	result := store.posts.FindOne(context.TODO(), filter)
	err = result.Decode(&product)
	return
}

func decode(cursor *mongo.Cursor) (products []*domain.Post, err error) {
	for cursor.Next(context.TODO()) {
		var product domain.Post
		err = cursor.Decode(&product)
		if err != nil {
			return
		}
		products = append(products, &product)
	}
	err = cursor.Err()
	return
}
