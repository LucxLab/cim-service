package handlers

import (
	"github.com/LucxLab/cim-service/internal/http/requests"
	"github.com/LucxLab/cim-service/internal/http/responses"
	"github.com/LucxLab/cim-service/internal/upload"
	"log"
	"net/http"
)

type Cdr struct {
	uploadUseCase upload.UseCase
}

func (h *Cdr) UploadFile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	uploadRequest, requestErr := requests.NewUpload(ctx, r)
	if requestErr != nil {
		log.Println(requestErr)
		responses.BadRequest(w)
		return
	}

	if validationErr := uploadRequest.Validate(); validationErr != nil {
		log.Println(validationErr)
		responses.BadRequest(w)
		return
	}

	fileCreation := uploadRequest.ToFileCreation()
	fileMetadata, useCaseErr := h.uploadUseCase.Process(ctx, fileCreation)
	if useCaseErr != nil {
		responses.UnexpectedError(w, useCaseErr)
		return
	}

	uploadResponse := responses.NewUpload(fileMetadata)
	if responseErr := responses.Json(w, http.StatusCreated, uploadResponse); responseErr != nil {
		responses.UnexpectedError(w, responseErr)
		return
	}
}

func NewCdr(uploadUseCase upload.UseCase) *Cdr {
	return &Cdr{
		uploadUseCase: uploadUseCase,
	}
}
