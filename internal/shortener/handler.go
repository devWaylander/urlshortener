package shortener

import (
	"context"
	"fmt"

	"github.com/devWhisper/urlshortener/gen/models"
	"github.com/devWhisper/urlshortener/gen/restapi/operations"
	"github.com/devWhisper/urlshortener/gen/restapi/operations/url"
	"github.com/go-openapi/runtime/middleware"
)

type UrlShortenerService interface {
	PutLongURL(ctx context.Context, longURL string) (shortURL string, err error)
	GetRedirectToLongURL(ctx context.Context, shortURL string) (longURL string, err error)
	GetRedirectionAnalytics(ctx context.Context, shortURL string) error
}

func Configure(api *operations.BackendCoreAPI, service UrlShortenerService) {
	api.URLPutLongHandler = url.PutLongHandlerFunc(func(params url.PutLongParams) middleware.Responder {
		if len(params.Data.LongURL) == 0 {
			return url.NewPutLongNotFound().WithPayload(&models.ErrorV1{Message: "Emprty req. URL must be provided"})
		}

		fmt.Println(params.Data.LongURL)
		token, err := service.PutLongURL(params.HTTPRequest.Context(), params.Data.LongURL)
		if err != nil {
			return url.NewPutLongNotFound().WithPayload(&models.ErrorV1{Message: err.Error()})
		}

		return url.NewPutLongOK().WithPayload(token)
	})

	api.URLGetLongHandler = url.GetLongHandlerFunc(func(params url.GetLongParams) middleware.Responder {
		if params.Token == "" {
			return url.NewGetLongNotFound().WithPayload(&models.ErrorV1{Message: "Emprty req. short URL must be provided"})
		}

		redirectToUrl, _ := service.GetRedirectToLongURL(params.HTTPRequest.Context(), params.Token)

		return url.NewGetLongMovedPermanently().WithLocation(redirectToUrl)
	})

	// api.URLGetAnalyticsHandler = url.GetAnalyticsHandlerFunc(func(params url.GetAnalyticsParams) middleware.Responder {

	// })
}
