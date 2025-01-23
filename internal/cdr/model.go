package cdr

type Upload struct {
	Id             string
	OrganizationId string
	UserId         string
	Title          string
	Location       string
	Status         UploadStatus
	CreatedAt      string
	UpdatedAt      string
}

type UploadStatus string

const (
	CreatedUploadStatus   UploadStatus = "created"
	SucceededUploadStatus              = "succeeded"
	FailedUploadStatus                 = "failed"
)
