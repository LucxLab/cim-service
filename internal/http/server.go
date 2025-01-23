package http

import (
	"fmt"
	"github.com/LucxLab/cim-service/internal/cdr"
	"github.com/LucxLab/cim-service/internal/minio"
	objStorageRepository "github.com/LucxLab/cim-service/internal/minio/repository"
	"github.com/LucxLab/cim-service/internal/mongo"
	dbRepository "github.com/LucxLab/cim-service/internal/mongo/repository"
	"github.com/LucxLab/cim-service/internal/rabbitmq"
	"github.com/LucxLab/cim-service/internal/rabbitmq/publisher"
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
	cdrDatabaseRepository := dbRepository.NewMongoCDR(mongoDatabase)

	// Object Storage Repositories
	cdrObjStorageRepository := objStorageRepository.NewMinioCdr(minioObjectStorage)

	// Publishers
	cdrPublisher := publisher.NewRabbitmqCdr(rabbitmqPublisher)

	// Services
	cdrService := cdr.NewService(cdrDatabaseRepository, cdrObjStorageRepository, cdrPublisher)

	router := NewRouter(cdrService)
	return &server{
		http: &http.Server{
			Addr:    address,
			Handler: router,
		},
	}
}
