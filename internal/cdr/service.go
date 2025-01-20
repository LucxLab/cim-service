package cdr

import (
	"fmt"
	"io"
	"strings"
)

type Service interface {
	UploadFile(file io.Reader, size int64, fileName string, companyId string) (*Upload, error)
}

type service struct {
	databaseRepository   DatabaseRepository
	objStorageRepository ObjectStorageRepository
}

func (s *service) UploadFile(file io.Reader, size int64, fileName string, companyId string) (*Upload, error) {
	uploadTitle := removeFileExtension(fileName)
	upload := &Upload{
		CompanyId: companyId,
		Title:     uploadTitle,
		Status:    CreatedUploadStatus,
	}

	err := s.databaseRepository.CreateUpload(upload)
	if err != nil {
		return nil, err
	}

	uploadLocation := fmt.Sprintf("%s/%s", companyId, upload.Id)
	err = s.objStorageRepository.SaveCdrFile(uploadLocation, size, file)
	if err != nil {
		return nil, err
	}

	upload.Location = uploadLocation
	upload.Status = SucceededUploadStatus
	err = s.databaseRepository.UpdateUpload(upload)
	if err != nil {
		return nil, err
	}
	return upload, nil
}

func NewService(databaseRepository DatabaseRepository, objStorageRepository ObjectStorageRepository) Service {
	return &service{databaseRepository: databaseRepository, objStorageRepository: objStorageRepository}
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
