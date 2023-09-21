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

func (r *repo) PutLongUrlAndToken(ctx context.Context, longURL string, shortURL string, isActive bool, expiresAt strfmt.DateTime) error {
	query := `INSERT INTO shorted_urls (long_url, token, is_active, expires_at) VALUES ($1, $2, $3, $4)`
	_, err := r.db.ExecContext(ctx, query, longURL, shortURL, isActive, expiresAt)

	return err
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

func (r *repo) GetCheckExistTokens(ctx context.Context) ([]string, error) {
	tokens := []string{}

	query := `SELECT token FROM shorted_urls`
	err := r.db.SelectContext(ctx, &tokens, query)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

func (r *repo) GetCheckExistLongUrl(ctx context.Context, longURL string) (*model.ShortUrl, error) {
	urlModel := model.ShortUrl{}
	query := `SELECT * FROM shorted_urls WHERE long_url = $1`

	existStatus := r.db.GetContext(ctx, &urlModel, query, longURL)
	if urlModel.LongUrl == longURL {
		return &urlModel, existStatus
	}

	return nil, nil
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
