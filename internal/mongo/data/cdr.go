package data

import "github.com/LucxLab/cim-service/internal/cdr"

type Upload struct {
	Status string `bson:"status"`
}

func FromUpload(upload *cdr.Upload) *Upload {
	return &Upload{
		Status: string(upload.Status),
	}
}
