package message

import "github.com/LucxLab/cim-service/internal/cdr"

type CdrFileUploaded struct {
	Id string `json:"id"`
}

func ToCdrFileUploaded(upload *cdr.Upload) *CdrFileUploaded {
	return &CdrFileUploaded{
		Id: upload.Id,
	}
}
