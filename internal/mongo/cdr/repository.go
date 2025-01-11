package cdr

import (
	cim "github.com/LucxLab/cim-service/internal"
)

const UploadCollection = "uploads"

type repository struct {
	db cim.Database
}

func (r *repository) CreateUpload(upload *cim.Upload) error {
	uploadData := &uploadData{
		Status: upload.Status,
	}

	identifier, err := r.db.InsertOne(UploadCollection, uploadData)
	if err != nil {
		return err
	}
	upload.Id = identifier
	return nil
}

func NewRepository(database cim.Database) cim.CDRRepository {
	return &repository{
		db: database,
	}
}
