package cdr

import (
	cim "github.com/LucxLab/cim-service/internal"
	"github.com/LucxLab/cim-service/internal/http/json"
	"github.com/gorilla/mux"
	"net/http"
)

const HandlerPath = "/cdr"
const UploadFilePath = ""

type Handler struct {
	*mux.Router
	uploadService cim.UploadService
}

func NewHandler(router *mux.Router, uploadService cim.UploadService) *Handler {
	cdrRouter := router.PathPrefix(HandlerPath).Subrouter()
	handler := &Handler{
		Router:        cdrRouter,
		uploadService: uploadService,
	}

	cdrRouter.HandleFunc(UploadFilePath, handler.uploadFile).Methods(http.MethodPost)
	return handler
}

func (h *Handler) uploadFile(w http.ResponseWriter, _ *http.Request) {
	upload, err := h.uploadService.InitializeFileUpload()
	if err != nil {
		json.WriteJSONResponse(w, http.StatusInternalServerError, map[string]string{
			"message": "Failed to initialize file upload",
		})
	}

	response := toUploadResponse(upload)
	json.WriteJSONResponse(w, http.StatusOK, response)
}
