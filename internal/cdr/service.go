package cdr

import (
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
	file io.Reader,
	size int64,
	fileName string,
	organizationId string,
	userId string,
) (*FileMetadata, error) {
	fileTitle := removeFileExtension(fileName)
	fileMetadata := &FileMetadata{
		OrganizationId:   organizationId,
		UserId:           userId,
		Title:            fileTitle,
		ProcessingStatus: CreatedFileProcessingStatus,
	}

	err := s.databaseRepository.CreateCdrFileMetadata(fileMetadata)
	if err != nil {
		return nil, err
	}

	uploadLocation := fmt.Sprintf("%s/%s", organizationId, fileMetadata.Id)
	err = s.objStorageRepository.SaveCdrFile(uploadLocation, size, file)
	if err != nil {
		//TODO: update file metadata status to UploadFailedFileProcessingStatus
		return nil, err
	}

	fileMetadata.Location = uploadLocation
	fileMetadata.ProcessingStatus = UploadSucceededFileProcessingStatus
	err = s.databaseRepository.UpdateCdrFileMetadata(fileMetadata)
	if err != nil {
		return nil, err
	}

	err = s.publisher.PublishCdrFileUploaded(fileMetadata.Id)
	if err != nil {
		return nil, err
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
