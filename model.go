package feedbacks

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Feedback struct {
	ID                        int                `pg:"id,pk" bson:"-" json:"-"`
	Id                        primitive.ObjectID `pg:"mongo_id" bson:"_id" json:"-"`
	Survey                    int                `pg:"survey" json:"survey"`
	Name                      string             `pg:"name" bson:",omitempty" json:"name"`
	Email                     string             `pg:"email" bson:",omitempty" json:"email"`
	Phone                     string             `pg:"phone" bson:",omitempty" json:"phone"`
	UserId                    string             `pg:"user_id" bson:"user_id,omitempty" json:"user_id"`
	Contacts                  string             `pg:"contacts" bson:",omitempty" json:"contacts"`
	Ts                        time.Time          `pg:"ts" bson:",omitempty" json:"ts"`
	Source                    string             `pg:"source" bson:",omitempty" json:"source"`
	Date                      string             `pg:"date" bson:",omitempty" json:"date"`
	Channel                   string             `pg:"channel" bson:",omitempty" json:"channel"`
	Nickname                  string             `pg:"nickname" bson:",omitempty" json:"nickname"`
	ClientStatus              string             `pg:"client_status" bson:"client_status,omitempty" json:"client_status"`
	Description               string             `pg:"description" bson:",omitempty" json:"description"`
	FeedbackCategory          string             `pg:"feedback_category" bson:"feedback_category,omitempty" json:"feedback_category"`
	ArticleCategory           string             `pg:"article_category" bson:"article_category,omitempty" json:"article_category"`
	QualityCategory           string             `pg:"quality_category" bson:"quality_category,omitempty" json:"quality_category"`
	AcceptanceQualityCategory string             `pg:"acceptance_quality_category" bson:"acceptance_quality_category,omitempty" json:"acceptance_quality_category"`
	Defect                    string             `pg:"defect" bson:"defect,omitempty" json:"defect"`
	Article                   string             `pg:"article" bson:",omitempty" json:"article"`
	StoreDate                 string             `pg:"store_date" bson:"store_date,omitempty" json:"store_date"`
	OrderId                   string             `pg:"order_id" bson:"order_id,omitempty" json:"order_id"`
	Receipt                   string             `pg:"receipt" bson:",omitempty" json:"receipt"`
	Store                     string             `pg:"store" bson:",omitempty" json:"store"`
	User                      string             `pg:"user" bson:",omitempty" json:"user"`
	Sent                      bool               `pg:"sent" bson:"sent" json:"sent"`
	MindboxSent               bool               `pg:"mindbox_sent" bson:"mindbox_sent" json:"mindbox_sent"`
	ArticlesCount             map[string]int     `pg:"-" bson:"-" json:"-"`
	Status                    string             `pg:"status" bson:",omitempty" json:"status"`
	Conclusion                string             `pg:"conclusion" bson:",omitempty" json:"conclusion"`
	SourceType                string             `pg:"source_type" bson:"source_type,omitempty" json:"source_type"`
	TransportCompany          string             `pg:"transport_company" bson:"transport_company,omitempty" json:"transport_company"`
	TransportCity             string             `pg:"transport_city" bson:"transport_city,omitempty" json:"transport_city"`
	Created                   time.Time          `pg:"created" bson:",omitempty" json:"created"`
	Updated                   time.Time          `pg:"updated" bson:"-" json:"-"`
}
