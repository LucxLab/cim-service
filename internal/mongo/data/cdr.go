package data

import (
	"github.com/LucxLab/cim-service/internal/cdr"
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

type CreateFileMetadata struct {
	OrganizationId   bson.ObjectID `bson:"organization_id"`
	UserId           bson.ObjectID `bson:"user_id"`
	Title            string        `bson:"title"`
	ProcessingStatus string        `bson:"processing_status"`
	CreatedAt        string        `bson:"created_at"`
	UpdatedAt        string        `bson:"updated_at"`
}

type UploadSucceeded struct {
	Location         string `bson:"location"`
	ProcessingStatus string `bson:"processing_status"`
	UpdatedAt        string `bson:"updated_at"`
}

type UploadFailed struct {
	ProcessingStatus string `bson:"processing_status"`
	UpdatedAt        string `bson:"updated_at"`
}

func ToCreateFileMetadata(fileMetadata *cdr.FileMetadata) (*CreateFileMetadata, error) {
	now := time.Now().UTC().Format(time.RFC3339)
	organizationObjectId, err := bson.ObjectIDFromHex(fileMetadata.OrganizationId)
	if err != nil {
		return nil, err
	}

	userObjectId, err := bson.ObjectIDFromHex(fileMetadata.UserId)
	if err != nil {
		return nil, err
	}

	return &CreateFileMetadata{
		OrganizationId:   organizationObjectId,
		UserId:           userObjectId,
		Title:            fileMetadata.Title,
		ProcessingStatus: fileMetadata.ProcessingStatus.String(),
		CreatedAt:        now,
		UpdatedAt:        now,
	}, nil
}

func ToUploadSucceeded(fileLocation string) *UploadSucceeded {
	now := time.Now().UTC().Format(time.RFC3339)
	return &UploadSucceeded{
		Location:         fileLocation,
		ProcessingStatus: cdr.UploadSucceededFileProcessingStatus.String(),
		UpdatedAt:        now,
	}
}

func ToUploadFailed() *UploadFailed {
	now := time.Now().UTC().Format(time.RFC3339)
	return &UploadFailed{
		ProcessingStatus: cdr.UploadFailedFileProcessingStatus.String(),
		UpdatedAt:        now,
	}
}
