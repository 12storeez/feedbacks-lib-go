package feedbacks

type Repository interface {
	FindOne(condition map[string]interface{}) (*Feedback, error)
	Update(filter, update map[string]interface{}) error
	InsertOne(feedback *Feedback) error
	CountFeedbacks(condition map[string]interface{}) (int, error)
}

// RepositoryPG...
type RepositoryPG interface {
	FindOneToday() (*Feedback, error)
	Update(fb *Feedback) error
	Insert(fb *Feedback) error
	CountFeedback(article string) (int, error)
}
