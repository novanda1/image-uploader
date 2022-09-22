package api

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/codedius/imagekit-go"
)

type UploadResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (a *API) Upload(w http.ResponseWriter, r *http.Request) {
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
		PublicKey:  a.config.IK.PubKey,
		PrivateKey: a.config.IK.PrivKey,
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

	return
}
