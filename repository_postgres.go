package feedbacks

import (
	"github.com/go-pg/pg/v10"
	"strconv"
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
	count, err := p.db.
		Model(&Feedback{}).
		Where("article LIKE concat('%', ?, '%')", article).
		Count()
	if err != nil {
		return 0, err
	}
	return count, err
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

func (p postgresRepository) SelectBy(id string) (*Feedback, error) {
	// если айдишник не можем преобразовать в число, скорее всего запрос идет по айдишнику монги
	// нужно для поддержки кнопок у одзывов для которых в апи слака сохранен айдишник монги
	pgID, err := strconv.Atoi(id)
	if err != nil {
		return p.selectBy(id)
	}

	result := &Feedback{
		ID: pgID,
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

func (p postgresRepository) selectBy(mongoID string) (*Feedback, error) {
	result := &Feedback{}

	if err := p.db.
		Model(result).
		Where("mongo_id = ?", mongoID).
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
		store = EXCLUDED.store, "user" = EXCLUDED.user, sent = EXCLUDED.sent, mindbox_sent = EXCLUDED.mindbox_sent,
		status = EXCLUDED.status, conclusion = EXCLUDED.conclusion, source_type = EXCLUDED.source_type, transport_company = EXCLUDED.transport_company,
		transport_city = EXCLUDED.transport_city, status_new = EXCLUDED.status_new, status_inwork = EXCLUDED.status_inwork, 
		status_done = EXCLUDED.status_done, status_err = EXCLUDED.status_err, updated = ?`, time.Now()).
		Insert(); err != nil {
		return err
	}

	return nil
}
