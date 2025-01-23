package cdr

import (
	"fmt"
	"io"
	"strings"
	"time"
)

type Service interface {
	UploadFile(file io.Reader, size int64, fileName string, organizationId string, userId string) (*Upload, error)
}

type service struct {
	databaseRepository   DatabaseRepository
	objStorageRepository ObjectStorageRepository
	publisher            Publisher
}

func (s *service) UploadFile(file io.Reader, size int64, fileName string, organizationId string, userId string) (*Upload, error) {
	uploadTitle := removeFileExtension(fileName)
	timeNow := time.Now().UTC().Format(time.RFC3339)
	upload := &Upload{
		OrganizationId: organizationId,
		UserId:         userId,
		Title:          uploadTitle,
		Status:         CreatedUploadStatus,
		CreatedAt:      timeNow,
		UpdatedAt:      timeNow,
	}

	err := s.databaseRepository.CreateUpload(upload)
	if err != nil {
		return nil, err
	}

	uploadLocation := fmt.Sprintf("%s/%s", organizationId, upload.Id)
	err = s.objStorageRepository.SaveCdrFile(uploadLocation, size, file)
	if err != nil {
		return nil, err
	}

	timeNow = time.Now().UTC().Format(time.RFC3339)
	upload.Location = uploadLocation
	upload.Status = SucceededUploadStatus
	upload.UpdatedAt = timeNow
	err = s.databaseRepository.UpdateUpload(upload)
	if err != nil {
		return nil, err
	}

	err = s.publisher.PublishUploadCreated(upload)
	if err != nil {
		return nil, err
	}
	return upload, nil
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
	subParts := strings.Split(fileName, ".")
	if len(subParts) == 1 {
		return fileName
	}
	return strings.Join(subParts[:len(subParts)-1], ".")
}
