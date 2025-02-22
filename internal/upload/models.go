package upload

import (
	"fmt"
	"io"
)

// FileCreation represents the data needed to create a CDR file.
type FileCreation struct {
	OrganizationId string
	UserId         string
	File           io.Reader
	FileSize       int64
	FileName       string
}

// GetStoringPath returns the file path where the file will be stored.
func (f *FileCreation) GetStoringPath(metadataId string) string {
	return fmt.Sprintf("%s/%s", f.OrganizationId, metadataId)
}

// FileMetadata represents the metadata of a CDR file uploaded.
type FileMetadata struct {
	Id     string
	Status FileUploadStatus
}

// FileUploadStatus represents the status of a CDR file upload.
type FileUploadStatus int

const (
	// CreatedFileUploadStatus represents the initial status of a CDR file when the upload starts, but it is not yet uploaded.
	CreatedFileUploadStatus FileUploadStatus = iota

	// SucceededFileUploadStatus represents the status of a CDR file successfully uploaded.
	SucceededFileUploadStatus

	// FailedFileUploadStatus represents the status of a CDR file that failed to upload.
	FailedFileUploadStatus
)

// String returns the string representation of the FileUploadStatus.
func (s FileUploadStatus) String() string {
	var fileUploadStatusNames = map[FileUploadStatus]string{
		CreatedFileUploadStatus:   "upload_created",
		SucceededFileUploadStatus: "upload_succeeded",
		FailedFileUploadStatus:    "upload_failed",
	}
	return fileUploadStatusNames[s]
}
