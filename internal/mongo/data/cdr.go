package data

import (
	"github.com/LucxLab/cim-service/internal/cdr"
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

type CreateCdrFileMetadata struct {
	OrganizationId   bson.ObjectID `bson:"organization_id"`
	UserId           bson.ObjectID `bson:"user_id"`
	Title            string        `bson:"title"`
	ProcessingStatus string        `bson:"processing_status"`
	CreatedAt        string        `bson:"created_at"`
	UpdatedAt        string        `bson:"updated_at"`
}

type UpdateCdrFileMetadata struct {
	Location         string `bson:"location"`
	ProcessingStatus string `bson:"processing_status"`
	UpdatedAt        string `bson:"updated_at"`
}

func ToCreateCdrFileMetadata(fileMetadata *cdr.FileMetadata) (*CreateCdrFileMetadata, error) {
	now := time.Now().UTC().Format(time.RFC3339)
	organizationObjectId, err := bson.ObjectIDFromHex(fileMetadata.OrganizationId)
	if err != nil {
		return nil, err
	}

	userObjectId, err := bson.ObjectIDFromHex(fileMetadata.UserId)
	if err != nil {
		return nil, err
	}

	return &CreateCdrFileMetadata{
		OrganizationId:   organizationObjectId,
		UserId:           userObjectId,
		Title:            fileMetadata.Title,
		ProcessingStatus: fileMetadata.ProcessingStatus.String(),
		CreatedAt:        now,
		UpdatedAt:        now,
	}, nil
}

func ToUpdateCdrFileMetadata(upload *cdr.FileMetadata) *UpdateCdrFileMetadata {
	now := time.Now().UTC().Format(time.RFC3339)
	return &UpdateCdrFileMetadata{
		Location:         upload.Location,
		ProcessingStatus: upload.ProcessingStatus.String(),
		UpdatedAt:        now,
	}
}
