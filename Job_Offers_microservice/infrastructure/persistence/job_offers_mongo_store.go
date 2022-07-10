package persistence

import (
	"context"
	"github.com/stojic19/XWS-TIM15/job_offers_microservice/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

const (
	DATABASE   = "job_offers"
	COLLECTION = "job_offers"
)

type JobOffersMongoStore struct {
	jobOffers *mongo.Collection
}

func NewJobOffersStore(client *mongo.Client) domain.JobOffersStore {
	jobOffers := client.Database(DATABASE).Collection(COLLECTION)
	return &JobOffersMongoStore{
		jobOffers: jobOffers,
	}
}

func (store *JobOffersMongoStore) GetAll() ([]*domain.JobOffer, error) {
	filter := bson.D{{}}
	return store.filter(filter)
}

func (store *JobOffersMongoStore) Get(id primitive.ObjectID) (*domain.JobOffer, error) {
	filter := bson.M{"_id": id}
	return store.filterOne(filter)
}

func (store *JobOffersMongoStore) GetSubscribed(userId string) ([]*domain.JobOffer, error) {
	filter := bson.M{"subscribers.id": userId}
	return store.filter(filter)
}

func (store *JobOffersMongoStore) GetRecommended(userId string, skills []string) ([]*domain.JobOffer, error) {
	filter := bson.M{"subscribers.id": bson.M{"$ne": userId}}
	offers, err := store.filter(filter)
	if err != nil {
		return nil, err
	}
	recommendedOffers := []*domain.JobOffer{}
	for _, jobOffer := range offers {
		requirements := strings.Split(jobOffer.Requirements, ",")
		requirementsMet := 0
		for _, requirement := range requirements {
			if Contains(skills, strings.Trim(requirement, " ")) {
				requirementsMet++
			}
		}
		if float64(requirementsMet)/float64(len(requirements)) > 0.75 {
			recommendedOffers = append(recommendedOffers, jobOffer)
		}
	}
	return recommendedOffers, nil
}

func (store *JobOffersMongoStore) Create(jobOffer *domain.JobOffer) error {
	jobOffer.Id = primitive.NewObjectID()
	result, err := store.jobOffers.InsertOne(context.TODO(), jobOffer)
	if err != nil {
		return err
	}
	jobOffer.Id = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (store *JobOffersMongoStore) Update(jobOffer *domain.JobOffer) error {
	_, err := store.jobOffers.UpdateOne(
		context.TODO(),
		bson.M{"_id": jobOffer.Id},
		bson.D{
			{"$set", bson.D{
				{"position", jobOffer.Position},
				{"description", jobOffer.Description},
				{"requirements", jobOffer.Requirements}}},
		})
	if err != nil {
		return err
	}
	return nil
}

func (store *JobOffersMongoStore) Subscribe(jobOfferId primitive.ObjectID, user *domain.User) error {
	_, err := store.jobOffers.UpdateOne(
		context.TODO(),
		bson.M{"_id": jobOfferId},
		bson.D{
			{"$addToSet", bson.D{{"subscribers", user}}},
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (store *JobOffersMongoStore) Unsubscribe(jobOfferId primitive.ObjectID, user *domain.User) error {
	_, err := store.jobOffers.UpdateOne(
		context.TODO(),
		bson.M{"_id": jobOfferId},
		bson.D{
			{"$pull", bson.D{{"subscribers", user}}},
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (store *JobOffersMongoStore) filter(filter interface{}) ([]*domain.JobOffer, error) {
	cursor, err := store.jobOffers.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())

	if err != nil {
		return nil, err
	}
	return decode(cursor)
}

func (store *JobOffersMongoStore) filterOne(filter interface{}) (post *domain.JobOffer, err error) {
	result := store.jobOffers.FindOne(context.TODO(), filter)
	err = result.Decode(&post)
	return
}

func decode(cursor *mongo.Cursor) (jobOffers []*domain.JobOffer, err error) {
	for cursor.Next(context.TODO()) {
		var jobOffer domain.JobOffer
		err = cursor.Decode(&jobOffer)
		if err != nil {
			return
		}
		jobOffers = append(jobOffers, &jobOffer)
	}
	err = cursor.Err()
	return
}

// Contains tells whether a contains x.
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
