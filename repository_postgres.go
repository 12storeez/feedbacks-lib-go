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

func (p postgresRepository) CountFeedbackBy(article string) (int, error) {

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
		Set(`
		mongo_id = EXCLUDED.mongo_id, survey = EXCLUDED.survey, name = EXCLUDED.name, email = EXCLUDED.email,
		phone = EXCLUDED.phone, user_id = EXCLUDED.user_id, contacts = EXCLUDED.contacts, ts = EXCLUDED.ts,
		source = EXCLUDED.source, date = EXCLUDED.date, channel = EXCLUDED.channel, nickname = EXCLUDED.nickname,
		client_status = EXCLUDED.client_status, description = EXCLUDED.description, feedback_category = EXCLUDED.feedback_category, 
		article_category = EXCLUDED.article_category, quality_category = EXCLUDED.quality_category, 
		acceptance_quality_category = EXCLUDED.acceptance_quality_category, defect = EXCLUDED.defect,
		article = EXCLUDED.article, store_date = EXCLUDED.store_date, order_id = EXCLUDED.order_id, receipt = EXCLUDED.receipt,
		store = EXCLUDED.store, user = EXCLUDED.user, sent = EXCLUDED.sent, mindbox_sent = EXCLUDED.mindbox_sent,
		status = EXCLUDED.status, conclusion = EXCLUDED.conclusion, source_type = EXCLUDED.source_type, transport_company = EXCLUDED.transport_company,
		transport_city = EXCLUDED.transport_city, updated = ?`, time.Now()).
		Insert(); err != nil {
		return err
	}

	return nil
}
