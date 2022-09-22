package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"

	"github.com/codedius/imagekit-go"
	"github.com/novanda1/image-uploader/conf"
)

func NewApi(config *conf.GlobalConfiguration) *chi.Mux {
	r := chi.NewRouter()

	// cors
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   strings.Split(config.API.ExternalURL, ","),
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
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

	r.Mount("/v1", V1(config))

	return r
}

func V1(config *conf.GlobalConfiguration) http.Handler {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API v1"))
	})
	r.Mount("/image", Image(config))
	return r
}

type UploadParams struct {
	File string `json:"file"`
	Name string `json:"name"`
}

func Image(config *conf.GlobalConfiguration) http.Handler {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	r.Post("/upload", func(w http.ResponseWriter, r *http.Request) {
		var params UploadParams

		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			w.Write([]byte("failed to parse body"))
		}

		opts := imagekit.Options{
			PublicKey:  config.IK.PubKey,
			PrivateKey: config.IK.PrivKey,
		}

		ik, err := imagekit.NewClient(&opts)
		if err != nil {
			w.Write([]byte("failed to setup imagekit"))
		}

		ur := imagekit.UploadRequest{
			File:              params.File, // []byte OR *url.URL OR url.URL OR base64 string
			FileName:          params.Name,
			UseUniqueFileName: false,
			Tags:              []string{},
			Folder:            "/image-uploader",
			IsPrivateFile:     false,
			CustomCoordinates: "",
			ResponseFields:    nil,
		}

		resp, err := ik.Upload.ServerUpload(context.Background(), &ur)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(resp)
	})

	r.Get("/{image_id}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("handle get image by id"))
	})

	return r
}
