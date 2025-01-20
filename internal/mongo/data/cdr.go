package data

import (
	"github.com/LucxLab/cim-service/internal/cdr"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type CreateUpload struct {
	CompanyId bson.ObjectID `bson:"company_id"`
	Title     string        `bson:"title"`
	Status    string        `bson:"status"`
}

type UpdateUpload struct {
	Location string `bson:"location"`
	Status   string `bson:"status"`
}

func ToCreateUpload(upload *cdr.Upload) (*CreateUpload, error) {
	companyObjectId, err := bson.ObjectIDFromHex(upload.CompanyId)
	if err != nil {
		return nil, err
	}

	return &CreateUpload{
		CompanyId: companyObjectId,
		Title:     upload.Title,
		Status:    string(upload.Status),
	}, nil
}

func ToUpdateUpload(upload *cdr.Upload) *UpdateUpload {
	return &UpdateUpload{
		Location: upload.Location,
		Status:   string(upload.Status),
	}
}
