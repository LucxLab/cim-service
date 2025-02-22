package upload

import (
	"context"
	"github.com/LucxLab/cim-service/internal/publishers"
)

type useCase struct {
	databaseRepository      DatabaseRepository
	objectStorageRepository ObjectStorageRepository
	cdrPublisher            publishers.Cdr
}

func (u *useCase) Process(ctx context.Context, fileCreation *FileCreation) (*FileMetadata, error) {
	fileMetadata := &FileMetadata{}

	// Create file metadata in the database.
	metadataId, databaseErr := u.databaseRepository.CreateMetadata(fileCreation)
	if databaseErr != nil {
		return fileMetadata, databaseErr
	}
	fileMetadata.Id = metadataId

	// Save the file in the object storage.
	filePath := fileCreation.GetStoringPath(metadataId)
	if objStorageErr := u.objectStorageRepository.SaveCdrFile(filePath, fileCreation.FileSize, fileCreation.File); objStorageErr != nil {
		fileMetadata.Status = FailedFileUploadStatus

		// If the saving fails, update the file metadata with the upload failed status.
		if databaseErr := u.databaseRepository.UploadFailed(metadataId); databaseErr != nil {
			return fileMetadata, databaseErr
		}
		return fileMetadata, objStorageErr
	}
	fileMetadata.Status = SucceededFileUploadStatus

	// Update the file metadata with the upload succeeded status and the file location.
	if databaseErr := u.databaseRepository.UploadSucceeded(metadataId, filePath); databaseErr != nil {
		return fileMetadata, databaseErr
	}

	// Publish a command to process the CDR file that was successfully uploaded.
	if publishErr := u.cdrPublisher.PublishFileProcessCommand(ctx, metadataId); publishErr != nil {
		return fileMetadata, publishErr
	}
	return fileMetadata, nil
}

func NewUseCase(
	databaseRepository DatabaseRepository,
	objectStorageRepository ObjectStorageRepository,
	cdrPublisher publishers.Cdr,
) UseCase {
	return &useCase{
		databaseRepository:      databaseRepository,
		objectStorageRepository: objectStorageRepository,
		cdrPublisher:            cdrPublisher,
	}
}
