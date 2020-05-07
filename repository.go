package feedbacks

type Repository interface {
	FindOne(condition map[string]interface{}) (*Feedback, error)
	Update(filter, update map[string]interface{}) error
	InsertOne(feedback *Feedback) error
}
