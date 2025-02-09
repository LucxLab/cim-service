package response

import "github.com/LucxLab/cim-service/internal/cdr"

type CdrFileUploaded struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}

func ToCdrFileUploaded(upload *cdr.FileMetadata) *CdrFileUploaded {
	return &CdrFileUploaded{
		Id:     upload.Id,
		Status: upload.ProcessingStatus.String(),
	}
}
