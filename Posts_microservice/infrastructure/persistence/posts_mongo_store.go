package persistence

import (
	"context"
	"errors"
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

func (store *PostsMongoStore) Get(id primitive.ObjectID) (*domain.Post, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *PostsMongoStore) GetFromUser(id string) ([]*domain.Post, error) {
	filter := bson.M{"owner.id": id}
	return store.filter(filter)
}

func (store *PostsMongoStore) GetFromUsers(ids []string) ([]*domain.Post, error) {
	filter := bson.M{"owner.id": bson.M{"$in": ids}}
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

func (store *PostsMongoStore) LikePost(postId primitive.ObjectID, user *domain.User) error {
	result, err := store.posts.UpdateOne(
		context.TODO(),
		bson.M{"_id": postId},
		bson.D{
			{"$addToSet", bson.D{{"likes", user}}},
		},
	)
	if err != nil {
		return err
	}
	if result.MatchedCount > 1 {
		return errors.New("one document should've been updated")
	}
	return nil
}

func (store *PostsMongoStore) DislikePost(postId primitive.ObjectID, user *domain.User) error {
	result, err := store.posts.UpdateOne(
		context.TODO(),
		bson.M{"_id": postId},
		bson.D{
			{"$addToSet", bson.D{{"dislikes", user}}},
		},
	)
	if err != nil {
		return err
	}
	if result.MatchedCount > 1 {
		return errors.New("one document should've been updated")
	}
	return nil
}

func (store *PostsMongoStore) RemoveLike(postId primitive.ObjectID, user *domain.User) error {
	result, err := store.posts.UpdateOne(
		context.TODO(),
		bson.M{"_id": postId},
		bson.D{
			{"$pull", bson.D{{"likes", user}}},
		},
	)
	if err != nil {
		return err
	}
	if result.MatchedCount > 1 {
		return errors.New("one document should've been updated")
	}
	return nil
}

func (store *PostsMongoStore) RemoveDislike(postId primitive.ObjectID, user *domain.User) error {
	result, err := store.posts.UpdateOne(
		context.TODO(),
		bson.M{"_id": postId},
		bson.D{
			{"$pull", bson.D{{"dislikes", user}}},
		},
	)
	if err != nil {
		return err
	}
	if result.MatchedCount > 1 {
		return errors.New("one document should've been updated")
	}
	return nil
}

func (store *PostsMongoStore) CreateComment(postId primitive.ObjectID, comment *domain.Comment) error {
	result, err := store.posts.UpdateOne(
		context.TODO(),
		bson.M{"_id": postId},
		bson.D{
			{"$push", bson.D{{"comments", comment}}},
		},
	)
	if err != nil {
		return err
	}
	if result.MatchedCount != 1 {
		return errors.New("one document should've been updated")
	}
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

func (store *PostsMongoStore) filterOne(filter interface{}) (post *domain.Post, err error) {
	result := store.posts.FindOne(context.TODO(), filter)
	err = result.Decode(&post)
	return
}

func decode(cursor *mongo.Cursor) (posts []*domain.Post, err error) {
	for cursor.Next(context.TODO()) {
		var post domain.Post
		err = cursor.Decode(&post)
		if err != nil {
			return
		}
		posts = append(posts, &post)
	}
	err = cursor.Err()
	return
}
