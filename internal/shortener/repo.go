package shortener

import (
	"context"

	"github.com/devWhisper/urlshortener/internal/shortener/model"
	"github.com/go-openapi/strfmt"
	"github.com/jmoiron/sqlx"
)

type repo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) *repo {
	return &repo{db: db}
}

func (r *repo) PutLongUrlAndToken(ctx context.Context, longURL string, shortURL string, isActive bool, expiresAt strfmt.DateTime) (error, bool) {
	urlModel := model.ShortUrl{}

	checkQuery := `SELECT token FROM shorted_urls WHERE token = $1`
	existStatus := r.db.GetContext(ctx, &urlModel, checkQuery, shortURL)
	if urlModel.Token != "" {
		return existStatus, true
	}

	query := `INSERT INTO shorted_urls (long_url, token, is_active, expires_at) VALUES ($1, $2, $3, $4)`
	_, err := r.db.ExecContext(ctx, query, longURL, shortURL, isActive, expiresAt)

	return err, false
}

func (r *repo) GetShortedUrlByToken(ctx context.Context, token string) (*model.ShortUrl, error) {
	urlModel := model.ShortUrl{}
	query := `SELECT * FROM shorted_urls WHERE token = $1`

	err := r.db.GetContext(ctx, &urlModel, query, token)
	if err != nil {
		return nil, err
	}

	return &urlModel, nil
}

func (r *repo) GetCheckExistLongUrl(ctx context.Context, longURL string) (*model.ShortUrl, error) {
	urlModel := model.ShortUrl{}
	query := `SELECT * FROM shorted_urls WHERE long_url = $1`

	err := r.db.GetContext(ctx, &urlModel, query, longURL)
	if err != nil {
		return nil, err
	}
	if urlModel.LongUrl == longURL {
		return nil, nil
	}

	return &urlModel, nil
}

func (r *repo) UpdateActiveShortedUrlStatus(ctx context.Context, token string, isActive bool) error {
	urlModel := model.ShortUrl{}
	query := `UPDATE shorted_urls SET is_active = $1 WHERE token = $2`

	err := r.db.GetContext(ctx, &urlModel, query, isActive, token)
	if err != nil {
		return err
	}

	return nil
}
