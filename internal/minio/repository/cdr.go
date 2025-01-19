package repository

import (
	"context"
	"github.com/LucxLab/cim-service/internal/cdr"
	"github.com/LucxLab/cim-service/internal/minio"
	mn "github.com/minio/minio-go/v7"
	"io"
)

type minioCdr struct {
	storage *minio.ObjectStorage
}

func (m *minioCdr) SaveObject(bucketName string, objectName string, objectSize int64, reader io.Reader) error {
	_, err := m.storage.Client.PutObject(context.TODO(), bucketName, objectName, reader, objectSize, mn.PutObjectOptions{})
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
