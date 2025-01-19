package minio

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const defaultHost = "localhost:9000"

type ObjectStorage struct {
	Client *minio.Client
}

func NewStorage() *ObjectStorage {
	client, err := minio.New(defaultHost, &minio.Options{
		Creds:  credentials.NewStaticV4("minioadmin", "minioadmin", ""),
		Secure: false,
	})
	if err != nil {
		panic(err) //TODO: handle error
	}

	return &ObjectStorage{
		Client: client,
	}
}
