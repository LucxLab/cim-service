package responses

import "github.com/LucxLab/cim-service/internal/upload"

type Upload struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}

func NewUpload(metadata *upload.FileMetadata) *Upload {
	return &Upload{
		Id:     metadata.Id,
		Status: metadata.Status.String(),
	}
}
