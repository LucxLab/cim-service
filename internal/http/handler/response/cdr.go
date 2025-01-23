package response

import "github.com/LucxLab/cim-service/internal/cdr"

type CreateUpload struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}

func ToCreatedUpload(upload *cdr.Upload) *CreateUpload {
	return &CreateUpload{
		Id:     upload.Id,
		Status: upload.Status.String(),
	}
}
