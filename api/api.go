package api

import (
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	"github.com/sirupsen/logrus"

	"github.com/novanda1/image-uploader/conf"
)

// API is the main REST API
type API struct {
	handler http.Handler
	config  *conf.GlobalConfiguration
	version string
}

func (a *API) ListenAndServe(hostAndPort string) {
	log := logrus.WithField("component", "api")
	server := &http.Server{
		Addr:              hostAndPort,
		Handler:           a.handler,
		ReadHeaderTimeout: 2 * time.Second, // to mitigate a Slowloris attack
	}

	done := make(chan struct{})
	defer close(done)

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.WithError(err).Fatal("http server listen failed")
	}
}

func NewApi(config *conf.GlobalConfiguration) *API {
	api := &API{config: config, version: "1"}
	r := chi.NewRouter()

	// cors
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: strings.Split(config.API.ExternalURL, ","),
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(httprate.LimitByIP(100, 1*time.Minute))

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Route("/image", func(r chi.Router) {
			r.Post("/", api.Upload)
		})
	})

	api.handler = r
	return api
}
