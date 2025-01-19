package cdr

type Upload struct {
	Id     string
	Status UploadStatus
}

type UploadStatus string

const (
	CreatedUploadStatus   UploadStatus = "created"
	SucceededUploadStatus              = "succeeded"
	FailedUploadStatus                 = "failed"
)
