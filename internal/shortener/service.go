package shortener

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/devWhisper/urlshortener/internal/shortener/model"
	"github.com/go-openapi/strfmt"
	"golang.org/x/exp/slices"
)

type Repo interface {
	PutLongUrlAndToken(ctx context.Context, longURL string, shortURL string, isActive bool, expiresAt strfmt.DateTime) error
	GetCheckExistLongUrl(ctx context.Context, longURL string) (*model.ShortUrl, error)
	GetCheckExistTokens(ctx context.Context) ([]string, error)
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

func generateToken(n int) string {
	var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	token := make([]rune, n)

	for i := range token {
		token[i] = letters[rand.Intn(len(letters))]
	}

	return string(token)
}

func (s *service) PutLongURL(ctx context.Context, longURL string) (shortURL string, err error) {
	checkExistToken, _ := s.repo.GetCheckExistLongUrl(ctx, longURL)
	if checkExistToken != nil {
		return checkExistToken.Token, nil
	}

	// TODO: подумать
	tokens, _ := s.repo.GetCheckExistTokens(ctx)
	token := generateToken(15)

	for slices.Contains(tokens, token) {
		token = generateToken(15)
	}

	expiresAt := strfmt.DateTime(time.Now().AddDate(0, 0, 7))
	err = s.repo.PutLongUrlAndToken(ctx, longURL, token, true, expiresAt)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *service) GetRedirectToLongURL(ctx context.Context, shortURL string) error {
	fmt.Println("")

	return nil
}

func (s *service) GetRedirectionAnalytics(ctx context.Context, shortURL string) error {
	fmt.Println("")

	return nil
}
