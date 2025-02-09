package handler

import (
	"github.com/LucxLab/cim-service/internal/cdr"
	"github.com/LucxLab/cim-service/internal/http/handler/response"
	"net/http"
)

type cdrHandler struct {
	service cdr.Service
}

func (h *cdrHandler) UploadCdrFile(w http.ResponseWriter, r *http.Request) {
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

	fileMetadata, err := h.service.UploadCdrFile(file, headers.Size, headers.Filename, organizationId, userId)
	if err != nil {
		response.UnexpectedError(w, err)
		return
	}

	cdrFileUploadedResponse := response.ToCdrFileUploaded(fileMetadata)
	err = response.Json(w, http.StatusCreated, cdrFileUploadedResponse)
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
