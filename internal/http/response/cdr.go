package response

import "github.com/LucxLab/cim-service/internal/cdr"

type CreateUpload struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}

func FromUpload(upload *cdr.Upload) *CreateUpload {
	return &CreateUpload{
		Id:     upload.Id,
		Status: string(upload.Status),
	}
}
