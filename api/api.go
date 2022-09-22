package api

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/codedius/imagekit-go"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"

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

type UploadResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Image(config *conf.GlobalConfiguration) http.Handler {
	r := chi.NewRouter()

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(32 << 20)
		if err != nil {
			var resp UploadResponse
			resp.Message = "Payload not valid"
			resp.Status = "error"
			resp.Data = nil
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(resp)

			return
		}

		f, _, err := r.FormFile("file")
		if err != nil {
			var resp UploadResponse
			resp.Message = "file not valid"
			resp.Status = "error"
			resp.Data = nil
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(resp)

			return
		}

		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, f); err != nil {
			var resp UploadResponse
			resp.Message = err.Error()
			resp.Status = "error"
			resp.Data = nil
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(resp)

			return
		}

		opts := imagekit.Options{
			PublicKey:  config.IK.PubKey,
			PrivateKey: config.IK.PrivKey,
		}

		ik, err := imagekit.NewClient(&opts)
		if err != nil {
			var resp UploadResponse
			resp.Message = "Server error"
			resp.Status = "error"
			resp.Data = nil
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(resp)
			return
		}

		ur := imagekit.UploadRequest{
			File:              buf.Bytes(), // []byte OR *url.URL OR url.URL OR base64 string
			FileName:          r.Form.Get("name"),
			UseUniqueFileName: false,
			Tags:              []string{},
			Folder:            "/image-uploader",
			IsPrivateFile:     false,
			CustomCoordinates: "",
			ResponseFields:    nil,
		}

		var resp UploadResponse

		ikresp, err := ik.Upload.ServerUpload(context.Background(), &ur)
		if err != nil {
			resp.Message = err.Error()
			resp.Status = "error"
			resp.Data = nil

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(resp)

			return
		}

		resp.Message = "Upload Successfully"
		resp.Status = "success"
		resp.Data = ikresp

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(resp)
	})

	return r
}
