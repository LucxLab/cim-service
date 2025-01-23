package handler

import (
	"github.com/LucxLab/cim-service/internal/cdr"
	"github.com/LucxLab/cim-service/internal/http/handler/response"
	"net/http"
)

type cdrHandler struct {
	service cdr.Service
}

func (h *cdrHandler) UploadFile(w http.ResponseWriter, r *http.Request) {
	organizationId := r.PathValue("organization_id")
	if organizationId == "" {
		response.BadRequest(w)
		return
	}

	userId := r.Header.Get("X-User-Id")
	if userId == "" {
		response.BadRequest(w)
		return
	}

	file, headers, err := r.FormFile("file")
	if err != nil {
		response.BadRequest(w)
		return
	}

	upload, err := h.service.UploadFile(file, headers.Size, headers.Filename, organizationId, userId)
	if err != nil {
		response.UnexpectedError(w, err)
		return
	}

	createdUploadResponse := response.ToCreatedUpload(upload)
	err = response.Json(w, http.StatusCreated, createdUploadResponse)
	if err != nil {
		response.UnexpectedError(w, err)
		return
	}
}

func NewCdrHandler(service cdr.Service) cdr.Handler {
	return &cdrHandler{
		service: service,
	}
}
