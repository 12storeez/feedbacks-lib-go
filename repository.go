package feedbacks

import "time"

type Repository interface {
	FindOne(condition map[string]interface{}) (*Feedback, error)
	Update(filter, update map[string]interface{}) error
	InsertOne(feedback *Feedback) error
	CountFeedbacks(condition map[string]interface{}) (int, error)
}

// RepositoryPG...
type RepositoryPG interface {
	FindOneNoSentToSlack(afterBorderTime time.Time) (*Feedback, error)
	Update(fb *Feedback) error
	Insert(fb *Feedback) error
	CountByArticle(article string) (int, error)
	FindByID(id int) (*Feedback, error)
}
