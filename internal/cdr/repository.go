package cdr

import "io"

type DatabaseRepository interface {
	CreateUpload(upload *Upload) error
	UpdateUpload(upload *Upload) error
}

type ObjectStorageRepository interface {
	SaveCdrFile(objectName string, objectSize int64, reader io.Reader) error
}
