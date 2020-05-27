package feedbacks

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	FEEDBACKS_COLLECTION = "survey_feedbacks"
)

var (
	ErrNoFeedbacks = mongo.ErrNoDocuments
)

type Mongo struct {
	Client   *mongo.Client
	Database *mongo.Database
}

type mongoRepository struct {
	mongo *Mongo
}

func NewMongoRepository(mongo *Mongo) Repository {
	return &mongoRepository{
		mongo: mongo,
	}
}

func (r mongoRepository) FindOne(condition map[string]interface{}) (*Feedback, error) {
	fb := &Feedback{}
	err := r.mongo.Database.Collection(FEEDBACKS_COLLECTION).FindOne(context.TODO(), condition).Decode(fb)
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
	_, err := r.mongo.Database.Collection(FEEDBACKS_COLLECTION).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return fmt.Errorf("r.UpdateSent: UpdateOne: %v", err)
	}
	return nil
}

func (r mongoRepository) InsertOne(feedback *Feedback) error {
	_, err := r.mongo.Database.Collection(FEEDBACKS_COLLECTION).InsertOne(context.TODO(), feedback)
	if err != nil {
		return fmt.Errorf("r.InsertOne: InsertOne: %v", err)
	}

	return nil
}
