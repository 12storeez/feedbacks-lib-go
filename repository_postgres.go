package feedbacks

import (
	"github.com/go-pg/pg/v10"
	"time"
)

type postgres struct {
	db *pg.DB
}

// NewPostgresRepository...
func NewPostgresRepository(db *pg.DB) RepositoryPG {
	return &postgres{
		db: db,
	}
}

// FindOneToday...
// Возвращает фидбек, который сохранили сегодня через форму
// но еще не отправленный в slack
func (p postgres) FindOneToday() (*Feedback, error) {
	var result Feedback

	if err := p.db.
		Model(&result).
		Where("ts >= ?", getToday()).
		Where("sent = ?", false).
		Limit(1).
		Select(); err != nil {
		return nil, err
	}

	return &result, nil
}

// Update...
func (p postgres) Update(fb *Feedback) error {
	if _, err := p.db.Model(&fb).UpdateNotZero(); err != nil {
		return err
	}
	return nil
}

// Insert...
func (p postgres) Insert(fb *Feedback) error {
	if _, err := p.db.Model(&fb).Insert(); err != nil {
		return err
	}
	return nil
}

// CountFeedback...
func (p postgres) CountFeedback(article string) (int, error) {
	count, err := p.db.
		Model(&Feedback{}).
		Where("article = ?", article).
		Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func getToday() time.Time {
	t := time.Now()
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}
