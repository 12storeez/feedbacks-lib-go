package feedbacks

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Feedback struct {
	Id               primitive.ObjectID `bson:"_id"`
	Survey           int
	Name             string
	Email            string
	Phone            string
	UserId           string `bson:"user_id"`
	Contacts         string
	Ts               time.Time
	Source           string
	Date             string
	Channel          string
	Nickname         string
	ClientStatus     string `bson:"client_status"`
	Description      string
	FeedbackCategory string `bson:"feedback_category"`
	ArticleCategory  string `bson:"article_category"`
	QualityCategory  string `bson:"quality_category"`
	Article          string
	StoreDate        string `bson:"store_date"`
	OrderId          string `bson:"order_id"`
	Receipt          string
	Store            string
	User             string
	Sent             bool
	MindboxSent      bool `bson:"mindbox_sent"`
	ArticlesCount    map[string]int
	Status           string
	Created          time.Time
}
