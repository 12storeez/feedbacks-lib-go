package feedbacks

import (
	"github.com/go-pg/pg/v10"
	"time"
)

type postgres struct {
	db *pg.DB
}

// FindByID...
func (p postgres) FindBy(id int) (*Feedback, error) {
	var result Feedback

	if err := p.db.
		Model(&result).
		Where("id = ?", id).
		Select(); err != nil {
		return nil, err
	}

	return &result, nil
}

// NewPostgresRepository...
func NewPostgresRepository(db *pg.DB) RepositoryPG {
	return &postgres{
		db: db,
	}
}

// FindOneToday...
func (p postgres) FindOneNoSentToSlack(afterBorderTime time.Time) (*Feedback, error) {
	var result Feedback

	if err := p.db.
		Model(&result).
		Where("created >= ?", afterBorderTime).
		Where("sent = ?", false).
		Limit(1).
		Select(); err != nil {
		return nil, err
	}

	return &result, nil
}

// Update...
func (p postgres) Update(fb *Feedback) error {
	fb.Updated = time.Now()
	if _, err := p.db.
		Model(fb).
		WherePK().
		UpdateNotZero(); err != nil {
		return err
	}
	return nil
}

// Insert...
func (p postgres) Insert(fb *Feedback) error {
	if _, err := p.db.
		Model(fb).
		Insert(); err != nil {
		return err
	}
	return nil
}

// CountByArticle...
func (p postgres) CountByArticle(article string) (int, error) {
	count, err := p.db.
		Model(&Feedback{}).
		Where("article = ?", article).
		Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}
