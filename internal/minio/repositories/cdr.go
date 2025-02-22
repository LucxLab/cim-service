package repositories

import (
	"context"
	"github.com/LucxLab/cim-service/internal/minio"
	"github.com/LucxLab/cim-service/internal/repositories"
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

func NewMinioCdr(storage *minio.ObjectStorage) repositories.CdrObjectStorage {
	return &minioCdr{
		storage: storage,
	}
}
