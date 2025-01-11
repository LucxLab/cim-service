package upload

import (
	cim "github.com/LucxLab/cim-service/internal"
)

type uploadService struct {
	repository cim.CDRRepository
}

func (s *uploadService) InitializeFileUpload() (*cim.Upload, error) {
	upload := &cim.Upload{
		Status: cim.PendingUploadStatus,
	}

	err := s.repository.CreateUpload(upload)
	if err != nil {
		return nil, err
	}
	return upload, nil
}

func NewService(repository cim.CDRRepository) cim.UploadService {
	return &uploadService{
		repository: repository,
	}
}
