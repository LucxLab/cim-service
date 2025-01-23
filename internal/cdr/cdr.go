package cdr

import (
	"io"
	"net/http"
)

/*
// Types
*/

// Upload represents metadata of a CDR file upload.
type Upload struct {
	Id             string
	OrganizationId string
	UserId         string
	Title          string
	Location       string
	Status         UploadStatus
	CreatedAt      string
	UpdatedAt      string
}

// UploadStatus represents the status of a CDR file upload.
//
// Possible values are "created", "succeeded", and "failed".
type UploadStatus int

const (
	// CreatedUploadStatus represents the status of a CDR file upload when it is freshly created, but not yet uploaded.
	CreatedUploadStatus UploadStatus = iota

	// SucceededUploadStatus represents the status of a CDR file upload when it is successfully uploaded.
	SucceededUploadStatus

	// FailedUploadStatus represents the status of a CDR file upload when it fails to upload.
	FailedUploadStatus
)

// String returns the string representation of the UploadStatus.
func (s UploadStatus) String() string {
	var uploadStatusesNames = map[UploadStatus]string{
		CreatedUploadStatus:   "created",
		SucceededUploadStatus: "succeeded",
		FailedUploadStatus:    "failed",
	}

	return uploadStatusesNames[s]
}

/*
// Interfaces
*/

// Handler defines methods that a CDR handler must implement to handle HTTP requests.
type Handler interface {
	UploadFile(w http.ResponseWriter, r *http.Request)
}

// Publisher defines methods that a CDR publisher must implement to publish messages.
type Publisher interface {
	PublishUploadCreated(upload *Upload) error
}

// DatabaseRepository defines methods that a CDR repository must implement to interact with the database.
type DatabaseRepository interface {
	CreateUpload(upload *Upload) error
	UpdateUpload(upload *Upload) error
}

// ObjectStorageRepository defines methods that a CDR repository must implement to interact with the object storage.
type ObjectStorageRepository interface {
	SaveCdrFile(objectName string, objectSize int64, reader io.Reader) error
}
