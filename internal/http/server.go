package http

import (
	"fmt"
	"github.com/LucxLab/cim-service/internal/minio"
	objStorageRepository "github.com/LucxLab/cim-service/internal/minio/repositories"
	"github.com/LucxLab/cim-service/internal/mongo"
	dbRepository "github.com/LucxLab/cim-service/internal/mongo/repositories"
	"github.com/LucxLab/cim-service/internal/rabbitmq"
	"github.com/LucxLab/cim-service/internal/rabbitmq/publishers"
	"github.com/LucxLab/cim-service/internal/upload"
	"net/http"
)

type Server interface {
	Listen() error
	Close() error
}

type server struct {
	http *http.Server
}

func (s *server) Listen() error {
	fmt.Println("Server listening on", s.http.Addr)
	return s.http.ListenAndServe()
}

func (s *server) Close() error {
	return nil
}

func New(address string) Server {
	mongoDatabase := mongo.NewDatabase()
	minioObjectStorage := minio.NewStorage()
	rabbitmqPublisher := rabbitmq.NewPublisher()

	// Database Repositories
	cdrDatabaseRepository := dbRepository.NewMongoCdr(mongoDatabase)

	// Object Storage Repositories
	cdrObjStorageRepository := objStorageRepository.NewMinioCdr(minioObjectStorage)

	// Publishers
	cdrPublisher := publishers.NewRabbitmqCdr(rabbitmqPublisher)

	// UseCases
	uploadUseCase := upload.NewUseCase(cdrDatabaseRepository, cdrObjStorageRepository, cdrPublisher)

	router := NewRouter(uploadUseCase)
	return &server{
		http: &http.Server{
			Addr:    address,
			Handler: router,
		},
	}
}
