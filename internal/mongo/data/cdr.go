package data

import (
	"github.com/LucxLab/cim-service/internal/cdr"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type CreateUpload struct {
	OrganizationId bson.ObjectID `bson:"organization_id"`
	UserId         bson.ObjectID `bson:"user_id"`
	Title          string        `bson:"title"`
	Status         string        `bson:"status"`
	CreatedAt      string        `bson:"created_at"`
	UpdatedAt      string        `bson:"updated_at"`
}

type UpdateUpload struct {
	Location  string `bson:"location"`
	Status    string `bson:"status"`
	UpdatedAt string `bson:"updated_at"`
}

func ToCreateUpload(upload *cdr.Upload) (*CreateUpload, error) {
	organizationObjectId, err := bson.ObjectIDFromHex(upload.OrganizationId)
	if err != nil {
		return nil, err
	}

	userObjectId, err := bson.ObjectIDFromHex(upload.UserId)
	if err != nil {
		return nil, err
	}

	return &CreateUpload{
		OrganizationId: organizationObjectId,
		UserId:         userObjectId,
		Title:          upload.Title,
		Status:         string(upload.Status),
		CreatedAt:      upload.CreatedAt,
		UpdatedAt:      upload.UpdatedAt,
	}, nil
}

func ToUpdateUpload(upload *cdr.Upload) *UpdateUpload {
	return &UpdateUpload{
		Location:  upload.Location,
		Status:    string(upload.Status),
		UpdatedAt: upload.UpdatedAt,
	}
}
