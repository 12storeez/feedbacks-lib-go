package feedbacks

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Feedback struct {
	Id               primitive.ObjectID `bson:"_id" json:"-"`
	Survey           int                `json:"survey"`
	Name             string             `bson:",omitempty" json:"name"`
	Email            string             `bson:",omitempty" json:"email"`
	Phone            string             `bson:",omitempty" json:"phone"`
	UserId           string             `bson:"user_id,omitempty" json:"user_id"`
	Contacts         string             `bson:",omitempty" json:"contacts"`
	Ts               time.Time          `bson:",omitempty" json:"ts"`
	Source           string             `bson:",omitempty" json:"source"`
	Date             string             `bson:",omitempty" json:"date"`
	Channel          string             `bson:",omitempty" json:"channel"`
	Nickname         string             `bson:",omitempty" json:"nickname"`
	ClientStatus     string             `bson:"client_status,omitempty" json:"client_status"`
	Description      string             `bson:",omitempty" json:"description"`
	FeedbackCategory string             `bson:"feedback_category,omitempty" json:"feedback_category"`
	ArticleCategory  string             `bson:"article_category,omitempty" json:"article_category"`
	QualityCategory  string             `bson:"quality_category,omitempty" json:"quality_category"`
	Article          string             `bson:",omitempty" json:"article"`
	StoreDate        string             `bson:"store_date,omitempty" json:"store_date"`
	OrderId          string             `bson:"order_id,omitempty" json:"order_id"`
	Receipt          string             `bson:",omitempty" json:"receipt"`
	Store            string             `bson:",omitempty" json:"store"`
	User             string             `bson:",omitempty" json:"user"`
	Sent             bool               `bson:"sent" json:"sent"`
	MindboxSent      bool               `bson:"mindbox_sent" json:"mindbox_sent"`
	ArticlesCount    map[string]int     `bson:"-" json:"-"`
	Status           string             `bson:",omitempty" json:"status"`
	Conclusion       string             `bson:",omitempty" json:"conclusion"`
	SourceType       string             `bson:"source_type,omitempty" json:"source_type"`
	TransportCompany string             `bson:"transport_company,omitempty" json:"transport_company"`
	TransportCity    string             `bson:"Transport_city,omitempty" json:"Transport_city"`
	Created          time.Time          `bson:",omitempty" json:"created"`
}
