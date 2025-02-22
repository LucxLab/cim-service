package repositories

import (
	"github.com/LucxLab/cim-service/internal/upload"
)

// CdrDatabase defines methods that a CDR repository must implement to interact with the database.
type CdrDatabase interface {
	upload.DatabaseRepository
}

// CdrObjectStorage defines methods that a CDR repository must implement to interact with the object storage.
type CdrObjectStorage interface {
	upload.ObjectStorageRepository
}
