package cim

type (
	Database interface {
		InsertOne(collectionName string, object interface{}) (identifier string, err error)
		Close() error
	}

	CDRRepository interface {
		CreateUpload(upload *Upload) error
	}

	Upload struct {
		Id     string
		Status UploadStatus
	}

	UploadStatus string

	UploadService interface {
		InitializeFileUpload() (*Upload, error)
	}
)

const (
	PendingUploadStatus   UploadStatus = "pending"
	SucceededUploadStatus              = "succeeded"
	FailedUploadStatus                 = "failed"
)
