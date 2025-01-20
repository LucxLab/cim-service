package repository

import (
	"context"
	"github.com/LucxLab/cim-service/internal/cdr"
	"github.com/LucxLab/cim-service/internal/minio"
	mn "github.com/minio/minio-go/v7"
	"io"
)

const cdrFilesBucketName = "cdr-files"

type minioCdr struct {
	storage *minio.ObjectStorage
}

func (m *minioCdr) SaveCdrFile(objectName string, objectSize int64, reader io.Reader) error {
	_, err := m.storage.Client.PutObject(context.TODO(), cdrFilesBucketName, objectName, reader, objectSize, mn.PutObjectOptions{})
	if err != nil {
		return err
	}
	return nil
}

func NewMinioCdr(storage *minio.ObjectStorage) cdr.ObjectStorageRepository {
	return &minioCdr{
		storage: storage,
	}
}
