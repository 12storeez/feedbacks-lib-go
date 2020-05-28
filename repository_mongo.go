package feedbacks

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	FEEDBACKS_DEFAULT_COLLECTION = "survey_feedbacks"
)

var (
	ErrNoFeedbacks = mongo.ErrNoDocuments
)

type Mongo struct {
	Client   *mongo.Client
	Database *mongo.Database
}

type mongoRepository struct {
	mongo      *Mongo
	collection string
}

func NewMongoRepository(mongo *Mongo, collection string) Repository {
	if collection == "" {
		collection = FEEDBACKS_DEFAULT_COLLECTION
	}
	return &mongoRepository{
		mongo:      mongo,
		collection: collection,
	}
}

func (r mongoRepository) FindOne(condition map[string]interface{}) (*Feedback, error) {
	fb := &Feedback{}
	err := r.mongo.Database.Collection(r.collection).FindOne(context.TODO(), condition).Decode(fb)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return nil, ErrNoFeedbacks
		default:
			return nil, fmt.Errorf("find feedback error: %v", err)
		}
	}
	return fb, nil
}

func (r mongoRepository) Update(filter, update map[string]interface{}) error {
	_, err := r.mongo.Database.Collection(r.collection).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return fmt.Errorf("r.UpdateSent: UpdateOne: %v", err)
	}
	return nil
}

func (r mongoRepository) InsertOne(feedback *Feedback) error {
	_, err := r.mongo.Database.Collection(r.collection).InsertOne(context.TODO(), feedback)
	if err != nil {
		return fmt.Errorf("r.InsertOne: InsertOne: %v", err)
	}

	return nil
}

func (r mongoRepository) CountFeedbacks(condition map[string]interface{}) (int, error) {
	count, err := r.mongo.Database.Collection(r.collection).
		CountDocuments(context.Background(), condition, &options.CountOptions{})
	if err != nil {
		return 0, fmt.Errorf("CountFeedbacks: %v", err)
	}
	return int(count), nil
}
