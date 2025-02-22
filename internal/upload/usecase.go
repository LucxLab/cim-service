package upload

import (
	"context"
	"github.com/LucxLab/cim-service/internal/logger"
	"github.com/LucxLab/cim-service/internal/publishers"
)

type useCase struct {
	databaseRepository      DatabaseRepository
	objectStorageRepository ObjectStorageRepository
	cdrPublisher            publishers.Cdr
	logger                  logger.Logger
}

func (u *useCase) Process(ctx context.Context, fileCreation *FileCreation) (*FileMetadata, error) {
	fileMetadata := &FileMetadata{}
	u.logger.DebugWithData("Processing file upload.", newAdditionalData(fileCreation, fileMetadata))

	// Create file metadata in the database.
	metadataId, databaseErr := u.databaseRepository.CreateMetadata(fileCreation)
	if databaseErr != nil {
		u.logger.WarnWithData("Failed to create file metadata.", newAdditionalData(fileCreation, fileMetadata))
		return fileMetadata, databaseErr
	}
	fileMetadata.Id = metadataId

	// Save the file in the object storage.
	filePath := fileCreation.GetStoringPath(metadataId)
	if objStorageErr := u.objectStorageRepository.SaveCdrFile(filePath, fileCreation.FileSize, fileCreation.File); objStorageErr != nil {
		fileMetadata.Status = FailedFileUploadStatus

		// If the saving fails, update the file metadata with the upload failed status.
		if databaseErr := u.databaseRepository.UploadFailed(metadataId); databaseErr != nil {
			u.logger.WarnWithData("Failed to update file metadata with upload failed status.", newAdditionalData(fileCreation, fileMetadata))
			return fileMetadata, databaseErr
		}
		u.logger.InfoWithData("Failed to save file in object storage.", newAdditionalData(fileCreation, fileMetadata))
		return fileMetadata, objStorageErr
	}
	fileMetadata.Status = SucceededFileUploadStatus

	// Update the file metadata with the upload succeeded status and the file location.
	if databaseErr := u.databaseRepository.UploadSucceeded(metadataId, filePath); databaseErr != nil {
		u.logger.WarnWithData("Failed to update file metadata with upload succeeded status.", newAdditionalData(fileCreation, fileMetadata))
		return fileMetadata, databaseErr
	}

	// Publish a command to process the CDR file that was successfully uploaded.
	if publishErr := u.cdrPublisher.PublishFileProcessCommand(ctx, metadataId); publishErr != nil {
		u.logger.WarnWithData("Failed to publish file process command.", newAdditionalData(fileCreation, fileMetadata))
		return fileMetadata, publishErr
	}

	u.logger.DebugWithData("File upload processed successfully.", newAdditionalData(fileCreation, fileMetadata))
	return fileMetadata, nil
}

func NewUseCase(
	databaseRepository DatabaseRepository,
	objectStorageRepository ObjectStorageRepository,
	cdrPublisher publishers.Cdr,
	logger logger.Logger,
) UseCase {
	return &useCase{
		databaseRepository:      databaseRepository,
		objectStorageRepository: objectStorageRepository,
		cdrPublisher:            cdrPublisher,
		logger:                  logger,
	}
}

func newAdditionalData(fileCreation *FileCreation, fileMetadata *FileMetadata) map[string]interface{} {
	return map[string]interface{}{
		"organization_id": fileCreation.OrganizationId,
		"user_id":         fileCreation.UserId,
		"file_name":       fileCreation.FileName,
		"file_size":       fileCreation.FileSize,
		"metadata_id":     fileMetadata.Id,
		"status":          fileMetadata.Status.String(),
	}
}
