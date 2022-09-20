package api

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httprate"

	"github.com/novanda1/image-uploader/conf"
)

func NewApi(config *conf.GlobalConfiguration) *chi.Mux {
	r := chi.NewRouter()

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

	r.Mount("/v1", V1())

	return r
}

func V1() http.Handler {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API v1"))
	})
	r.Mount("/image", Image())
	return r
}

func Image() http.Handler {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	r.Post("/upload", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("handle upload"))
	})

	r.Get("/{image_id}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("handle get image by id"))
	})

	return r
}
