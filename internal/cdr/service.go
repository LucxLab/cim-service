package cdr

import (
	"context"
	"fmt"
	"io"
	"strings"
)

type service struct {
	databaseRepository   DatabaseRepository
	objStorageRepository ObjectStorageRepository
	publisher            Publisher
}

func (s *service) UploadCdrFile(
	ctx context.Context,
	file io.Reader,
	size int64,
	fileName string,
	organizationId string,
	userId string,
) (*FileMetadata, error) {
	fileTitle := removeFileExtension(fileName)
	fileMetadata := newFileMetadata(organizationId, userId, fileTitle)

	// Save the file metadata in the database.
	if databaseErr := s.databaseRepository.CreateFileMetadata(fileMetadata); databaseErr != nil {
		return nil, databaseErr
	}

	// Save the file in the object storage.
	uploadLocation := fmt.Sprintf("%s/%s", organizationId, fileMetadata.Id)
	if objStorageErr := s.objStorageRepository.SaveCdrFile(uploadLocation, size, file); objStorageErr != nil {
		// If the file upload fails, update the file metadata with the failed status.
		if databaseErr := s.databaseRepository.UploadFailed(fileMetadata.Id); databaseErr != nil {
			return nil, databaseErr
		}

		fileMetadata.ProcessingStatus = UploadFailedFileProcessingStatus
		return nil, objStorageErr
	}

	// Update the file metadata with the successful status and the file location.
	if databaseErr := s.databaseRepository.UploadSucceeded(fileMetadata.Id, uploadLocation); databaseErr != nil {
		return nil, databaseErr
	}
	fileMetadata.ProcessingStatus = UploadSucceededFileProcessingStatus

	// Publish a message that the CDR file was successfully uploaded.
	if publisherErr := s.publisher.PublishCdrFileUploaded(ctx, fileMetadata.Id); publisherErr != nil {
		return nil, publisherErr
	}
	return fileMetadata, nil
}

func NewService(
	databaseRepository DatabaseRepository,
	objStorageRepository ObjectStorageRepository,
	publisher Publisher,
) Service {
	return &service{
		databaseRepository:   databaseRepository,
		objStorageRepository: objStorageRepository,
		publisher:            publisher,
	}
}

// newFileMetadata represents the default initialisation of a CDR file metadata.
func newFileMetadata(
	organizationId string,
	userId string,
	title string,
) *FileMetadata {
	return &FileMetadata{
		OrganizationId:   organizationId,
		UserId:           userId,
		Title:            title,
		ProcessingStatus: CreatedFileProcessingStatus,
	}
}

// removeFileExtension removes the file extension from the file name
// and returns the file name without the extension.
//
// Example: removeFileExtension("2025-01-01.123456.csv") returns "2025-01-01.123456"
func removeFileExtension(fileName string) string {
	parts := strings.Split(fileName, ".")
	if len(parts) == 1 {
		return fileName
	}
	return strings.Join(parts[:len(parts)-1], ".")
}
