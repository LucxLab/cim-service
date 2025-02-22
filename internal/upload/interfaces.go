package upload

import (
	"context"
	"io"
)

// UseCase define the Process method to upload a CDR file.
type UseCase interface {
	// Process uploads a CDR file to the object storage and saves its metadata in the database.
	Process(ctx context.Context, file *FileCreation) (*FileMetadata, error)
}

// DatabaseRepository defines methods to interact with the database for the upload use case.
type DatabaseRepository interface {
	// CreateMetadata creates a new CDR file metadata in the database and returns its ID.
	CreateMetadata(fileCreation *FileCreation) (id string, err error)

	// UploadSucceeded updates the status of a CDR file metadata to UploadSucceededFileProcessingStatus and sets the file location.
	UploadSucceeded(id string, fileLocation string) error

	// UploadFailed updates the status of a CDR file metadata to UploadFailedFileProcessingStatus.
	UploadFailed(id string) error
}

// ObjectStorageRepository defines methods to interact with the object storage for the upload use case.
type ObjectStorageRepository interface {
	// SaveCdrFile saves a CDR file in the object storage.
	SaveCdrFile(objectName string, objectSize int64, reader io.Reader) error
}
