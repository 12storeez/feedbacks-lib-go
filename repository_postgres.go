package feedbacks

import (
	"github.com/go-pg/pg/v10"
	"time"
)

type postgresRepository struct {
	db *pg.DB
}

func NewPostgresRepository(db *pg.DB) RepositoryPG {
	return &postgresRepository{
		db: db,
	}
}

func (p postgresRepository) CountFeedbackBy(article int) (int, error) {

	return 1, nil
}

func (p postgresRepository) SelectOneForSlack() (*Feedback, error) {
	result := &Feedback{}

	if err := p.db.
		Model(result).
		Where("sent = ?", false).
		Limit(1).
		Select(); err != nil {
		switch err {
		case pg.ErrNoRows:
			return nil, ErrNoFeedbacks
		default:
			return nil, err
		}
	}

	return result, nil
}

func (p postgresRepository) SelectBy(id int) (*Feedback, error) {
	result := &Feedback{
		ID: id,
	}

	if err := p.db.
		Model(result).
		WherePK().
		Select(); err != nil {
		switch err {
		case pg.ErrNoRows:
			return nil, ErrNoFeedbacks
		default:
			return nil, err
		}
	}

	return result, nil
}

func (p postgresRepository) Upsert(fb *Feedback) error {
	if _, err := p.db.
		Model(fb).
		OnConflict("(id) DO UPDATE").
		Set("updated = ?, sent = EXCLUDED.sent", time.Now()).
		Insert(); err != nil {
		return err
	}

	return nil
}
