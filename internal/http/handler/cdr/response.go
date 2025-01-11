package cdr

import cim "github.com/LucxLab/cim-service/internal"

type uploadResponse struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}

func toUploadResponse(upload *cim.Upload) *uploadResponse {
	return &uploadResponse{
		Id:     upload.Id,
		Status: string(upload.Status),
	}
}
