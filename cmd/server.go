package cmd

import (
	"net/http"
	"os"
	"time"

	"github.com/devWhisper/urlshortener/gen/restapi"
	"github.com/devWhisper/urlshortener/gen/restapi/operations"
	"github.com/devWhisper/urlshortener/internal/shortener"
	"github.com/go-openapi/loads"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/kenshaw/snaker"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

func RunServer() {
	err := godotenv.Load()
	if err != nil {
		logrus.WithError(err).Panic("Unable to connect to DB")
	}

	host, err := os.Hostname()
	if err != nil {
		logrus.WithError(err).Fatal("Something went wrong")
	}

	logrus.SetLevel(logrus.InfoLevel)
	logrus.WithFields(logrus.Fields{
		"Host": host,
	}).Info("Service Startup")

	swaggerSpec, err := loads.Analyzed(restapi.FlatSwaggerJSON, "")
	if err != nil {
		panic(err)
	}

	api := operations.NewBackendCoreAPI(swaggerSpec)
	api.Logger = logrus.Printf

	db, err := newDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		logrus.WithError(err).Panic("Unable to connect to DB")
	}

	// jokerRepo := joke.NewRepo(db)
	// jokeService := joke.NewService(jokerRepo)
	// joke.Configure(api, jokeService)

	shortenerRepo := shortener.NewRepo(db)
	shortenerService := shortener.NewService(shortenerRepo)
	shortener.Configure(api, shortenerService)

	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.Port = 8080
	server.EnabledListeners = []string{"http"}

	c := cors.AllowAll()
	mux := http.NewServeMux()
	mux.Handle("/", api.Serve(nil))

	server.SetHandler(c.Handler(mux))

	if err := server.Serve(); err != nil {
		logrus.Panicln(err)
	}
}

func newDB(connectionString string) (*sqlx.DB, error) {
	d, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	d.SetMaxOpenConns(5)
	d.SetMaxIdleConns(5)
	d.SetConnMaxLifetime(5 * time.Minute)
	d.MapperFunc(snaker.CamelToSnake)

	return d, nil
}
