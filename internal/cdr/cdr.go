package cdr

import (
	"context"
	"io"
	"net/http"
)

/*
// Types
*/

// FileMetadata represents metadata of a CDR file.
type FileMetadata struct {
	Id               string
	OrganizationId   string
	UserId           string
	Title            string
	Location         string
	ProcessingStatus FileProcessingStatus
	CreatedAt        string
	UpdatedAt        string
}

// FileProcessingStatus represents the status of a CDR file processing.
type FileProcessingStatus int

const (
	// CreatedFileProcessingStatus represents the initial status of a CDR file when the upload starts, but it is not yet uploaded.
	CreatedFileProcessingStatus FileProcessingStatus = iota

	// UploadSucceededFileProcessingStatus represents the status of a CDR file successfully uploaded.
	UploadSucceededFileProcessingStatus

	// UploadFailedFileProcessingStatus represents the status of a CDR file that failed to upload.
	UploadFailedFileProcessingStatus
)

// String returns the string representation of the FileProcessingStatus.
func (s FileProcessingStatus) String() string {
	var fileProcessingStatusNames = map[FileProcessingStatus]string{
		CreatedFileProcessingStatus:         "upload_created",
		UploadSucceededFileProcessingStatus: "upload_succeeded",
		UploadFailedFileProcessingStatus:    "upload_failed",
	}
	return fileProcessingStatusNames[s]
}

/*
// Interfaces
*/

// Handler defines methods that a CDR handler must implement to handle HTTP requests.
type Handler interface {
	// UploadCdrFile handles the HTTP request to upload a CDR file.
	UploadCdrFile(writer http.ResponseWriter, request *http.Request)
}

// Service defines methods that a CDR service must implement to handle business logic.
type Service interface {
	// UploadCdrFile uploads a CDR file to the object storage and saves its metadata in the database.
	UploadCdrFile(ctx context.Context, file io.Reader, size int64, fileName string, organizationId string, userId string) (*FileMetadata, error)
}

// Publisher defines methods that a CDR publisher must implement to publish messages to a message broker.
type Publisher interface {
	// PublishCdrFileUploaded publishes a message to a message broker that a CDR file was uploaded.
	PublishCdrFileUploaded(ctx context.Context, id string) error
}

// DatabaseRepository defines methods that a CDR repository must implement to interact with the database.
type DatabaseRepository interface {
	// CreateFileMetadata creates a new CDR file metadata in the database.
	CreateFileMetadata(fileMetadata *FileMetadata) error

	// UploadSucceeded updates the status of a CDR file metadata to UploadSucceededFileProcessingStatus and sets the file location.
	UploadSucceeded(id string, fileLocation string) error

	// UploadFailed updates the status of a CDR file metadata to UploadFailedFileProcessingStatus.
	UploadFailed(id string) error
}

// ObjectStorageRepository defines methods that a CDR repository must implement to interact with the object storage.
type ObjectStorageRepository interface {
	// SaveCdrFile saves a CDR file in the object storage.
	SaveCdrFile(objectName string, objectSize int64, reader io.Reader) error
}
