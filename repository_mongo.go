package feedbacks

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	FEEDBACKS_DEFAULT_COLLECTION = "survey_feedbacks"
	FEEDBACKS_DEFAULT_DATABASE   = "portal"
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
	dbname     string
}

func NewMongoRepository(mongo *Mongo, dbname string, collection string) Repository {
	if collection == "" {
		collection = FEEDBACKS_DEFAULT_COLLECTION
	}

	if dbname == "" {
		dbname = FEEDBACKS_DEFAULT_DATABASE
	}
	return &mongoRepository{
		mongo:      mongo,
		collection: collection,
		dbname:     dbname,
	}
}

func (r mongoRepository) FindOne(condition map[string]interface{}) (*Feedback, error) {
	fb := &Feedback{}
	err := r.mongo.Client.Database(r.dbname).Collection(r.collection).FindOne(context.TODO(), condition).Decode(fb)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return nil, ErrNoFeedbacks
		default:
			return nil, err
		}
	}
	return fb, nil
}

func (r mongoRepository) Update(filter, update map[string]interface{}) error {
	_, err := r.mongo.Client.Database(r.dbname).Collection(r.collection).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r mongoRepository) InsertOne(feedback *Feedback) error {
	_, err := r.mongo.Client.Database(r.dbname).Collection(r.collection).InsertOne(context.TODO(), feedback)
	if err != nil {
		return err
	}

	return nil
}

func (r mongoRepository) CountFeedbacks(condition map[string]interface{}) (int, error) {
	count, err := r.mongo.Client.Database(r.dbname).Collection(r.collection).
		CountDocuments(context.Background(), condition, &options.CountOptions{})
	if err != nil {
		return 0, err
	}
	return int(count), nil
}
