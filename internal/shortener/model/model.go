package model

import (
	"github.com/devWhisper/urlshortener/gen/restapi/operations/url"
	"github.com/go-openapi/strfmt"
)

type ShortUrl struct {
	ID       int64  `json:"id"`
	Token    string `json:"token"`
	LongUrl  string `json:"long_url"`
	IsActive bool   `json:"is_active"`

	CreatedAt strfmt.DateTime `json:"created_at"`
	ExpiresAt strfmt.DateTime `json:"expires_at"`
}

func (u *ShortUrl) ToSwaggerShortUrl() *url.PutLongOK {
	result := &url.PutLongOK{Payload: u.Token}

	return result
}
