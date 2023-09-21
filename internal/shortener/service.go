package shortener

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/devWhisper/urlshortener/internal/shortener/model"
	"github.com/go-openapi/strfmt"
)

type Repo interface {
	PutLongUrlAndToken(ctx context.Context, longURL string, shortURL string, isActive bool, expiresAt strfmt.DateTime) (error, bool)
	GetCheckExistLongUrl(ctx context.Context, longURL string) (*model.ShortUrl, error)
	GetShortedUrlByToken(ctx context.Context, token string) (*model.ShortUrl, error)
	UpdateActiveShortedUrlStatus(ctx context.Context, token string, isActive bool) error
}

type service struct {
	repo Repo
}

func NewService(repo Repo) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) PutLongURL(ctx context.Context, longURL string) (shortURL string, err error) {
	// взять
	// сгенерить токен
	// положить длинный урл и токен в бд
	// вернуть токен хендлеру
	checkExistToken, err := s.repo.GetCheckExistLongUrl(ctx, longURL)
	if err != nil {
		return "", err
	}
	if checkExistToken != nil {
		return checkExistToken.Token, nil
	}

	b := make([]byte, 1)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	c := hex.EncodeToString(b)
	fmt.Println(c)

	return
}

func (s *service) GetRedirectToLongURL(ctx context.Context, shortURL string) error {
	fmt.Println("")

	return nil
}

func (s *service) GetRedirectionAnalytics(ctx context.Context, shortURL string) error {
	fmt.Println("")

	return nil
}
