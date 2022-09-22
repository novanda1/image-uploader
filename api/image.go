package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"strings"

	"github.com/codedius/imagekit-go"
)

type UploadResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func badRequest(w http.ResponseWriter, msg string) {
	var resp UploadResponse
	resp.Message = msg
	resp.Status = "error"
	resp.Data = nil

	sendJSON(w, http.StatusBadRequest, resp)
}

func (a *API) Upload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		badRequest(w, "failed to parse form")
		return
	}

	f, h, err := r.FormFile("file")
	if err != nil {
		badRequest(w, "file not valid")
		return
	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, f); err != nil {
		badRequest(w, err.Error())
		return
	}

	if strings.Contains(h.Header.Get("Content-Type"), "image") != true {
		badRequest(w, "file must be image")
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

		sendJSON(w, http.StatusInternalServerError, resp)
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

		sendJSON(w, http.StatusInternalServerError, resp)
		return
	}

	resp.Message = "Upload Successfully"
	resp.Status = "success"
	resp.Data = ikresp

	sendJSON(w, http.StatusCreated, resp)
	return
}
