package feedbacks

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	FEEDBACKS_COLLECTION = "survey_feedbacks"
)

var (
	ErrNoFeedbacks = mongo.ErrNoDocuments
)

type Mongo struct {
	client *mongo.Client
	DB     *mongo.Database
}

type mongoRepository struct {
	mongo *Mongo
}

func NewMongoRepository(mongo *Mongo) Repository {
	return &mongoRepository{
		mongo: mongo,
	}
}

func (r mongoRepository) FindOne(condition map[string]interface{}) (*Feedback, error)  {
	fb := &Feedback{}
	err := r.mongo.DB.Collection(FEEDBACKS_COLLECTION).FindOne(context.TODO(), condition).Decode(fb)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return nil, ErrNoFeedbacks
		default:
			return nil, errors.Errorf("find feedback error: %v", err)
		}
	}
	return fb, nil
}

func (r mongoRepository) Update(filter, update map[string]interface{}) error {
	_, err := r.mongo.DB.Collection(FEEDBACKS_COLLECTION).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return errors.Errorf("r.UpdateSent: UpdateOne: %v", err)
	}
	return nil
}
