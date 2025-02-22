package data

import (
	"github.com/LucxLab/cim-service/internal/upload"
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

type FileMetadataCreation struct {
	OrganizationId bson.ObjectID `bson:"organization_id"`
	UserId         bson.ObjectID `bson:"user_id"`
	Name           string        `bson:"name"`
	Status         string        `bson:"status"`
	CreatedAt      string        `bson:"created_at"`
	UpdatedAt      string        `bson:"updated_at"`
}

type FileUploadSucceeded struct {
	Path      string `bson:"path"`
	Status    string `bson:"status"`
	UpdatedAt string `bson:"updated_at"`
}

type FileUploadFailed struct {
	Status    string `bson:"status"`
	UpdatedAt string `bson:"updated_at"`
}

func NewFileMetadataCreation(fileCreation *upload.FileCreation) (*FileMetadataCreation, error) {
	now := time.Now().UTC().Format(time.RFC3339)

	organizationObjectId, err := bson.ObjectIDFromHex(fileCreation.OrganizationId)
	if err != nil {
		return nil, err
	}

	userObjectId, err := bson.ObjectIDFromHex(fileCreation.UserId)
	if err != nil {
		return nil, err
	}

	return &FileMetadataCreation{
		OrganizationId: organizationObjectId,
		UserId:         userObjectId,
		Name:           fileCreation.FileName,
		Status:         upload.CreatedFileUploadStatus.String(),
		CreatedAt:      now,
		UpdatedAt:      now,
	}, nil
}

func NewFileUploadSucceeded(fileLocation string) *FileUploadSucceeded {
	now := time.Now().UTC().Format(time.RFC3339)
	return &FileUploadSucceeded{
		Path:      fileLocation,
		Status:    upload.SucceededFileUploadStatus.String(),
		UpdatedAt: now,
	}
}

func NewFileUploadFailed() *FileUploadFailed {
	now := time.Now().UTC().Format(time.RFC3339)
	return &FileUploadFailed{
		Status:    upload.FailedFileUploadStatus.String(),
		UpdatedAt: now,
	}
}
