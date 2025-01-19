package cdr

import (
	"io"
)

type Service interface {
	UploadFile(file io.Reader, size int64) (*Upload, error)
}

type service struct {
	databaseRepository   DatabaseRepository
	objStorageRepository ObjectStorageRepository
}

func (s *service) UploadFile(file io.Reader, size int64) (*Upload, error) {
	upload := &Upload{
		Status: CreatedUploadStatus,
	}
	err := s.databaseRepository.CreateUpload(upload)
	if err != nil {
		return nil, err
	}

	err = s.objStorageRepository.SaveObject("cdr-uploads", upload.Id, size, file)
	if err != nil {
		return nil, err
	}

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
