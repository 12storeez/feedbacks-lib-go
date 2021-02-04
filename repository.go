package feedbacks

type Repository interface {
	FindOne(condition map[string]interface{}) (*Feedback, error)
	Update(filter, update map[string]interface{}) error
	InsertOne(feedback *Feedback) error
	CountFeedbacks(condition map[string]interface{}) (int, error)
}

type RepositoryPG interface {
	SelectOneForSlack() (*Feedback, error)
	SelectBy(id string) (*Feedback, error)
	Upsert(fb *Feedback) error
	CountFeedbackBy(article string) (int, error)
}
