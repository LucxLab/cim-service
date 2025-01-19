package handler

import (
	"github.com/LucxLab/cim-service/internal/cdr"
	"github.com/LucxLab/cim-service/internal/http/json"
	"github.com/LucxLab/cim-service/internal/http/response"
	"net/http"
)

type handler struct {
	service cdr.Service
}

func (h *handler) UploadFile(w http.ResponseWriter, r *http.Request) {
	file, headers, err := r.FormFile("file")
	if err != nil {
		json.BadRequest(w)
		return
	}

	upload, err := h.service.UploadFile(file, headers.Size)
	if err != nil {
		json.UnexpectedError(w, err)
		return
	}

	uploadResponse := response.FromUpload(upload)
	err = json.Response(w, http.StatusCreated, uploadResponse)
	if err != nil {
		json.UnexpectedError(w, err)
		return
	}
}

func NewCdrHandler(service cdr.Service) cdr.Handler {
	return &handler{
		service: service,
	}
}
