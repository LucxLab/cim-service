package http

import (
	"fmt"
	"github.com/LucxLab/cim-service/internal/logger"
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
	http          *http.Server
	logger        logger.Logger
	mongoDatabase *mongo.Database
}

func (s *server) Listen() error {
	fmt.Println("Server listening on", s.http.Addr)
	return s.http.ListenAndServe()
}

func (s *server) Close() error {
	if loggerErr := s.logger.Close(); loggerErr != nil {
		return loggerErr
	}
	if mongoErr := s.mongoDatabase.Close(); mongoErr != nil {
		return mongoErr
	}
	return nil
}

func New(address string) Server {
	serverLogger, loggerErr := logger.New(true)
	if loggerErr != nil {
		panic(loggerErr)
	}

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
	uploadUseCase := upload.NewUseCase(cdrDatabaseRepository, cdrObjStorageRepository, cdrPublisher, serverLogger)

	router := NewRouter(uploadUseCase)
	return &server{
		http: &http.Server{
			Addr:    address,
			Handler: router,
		},
		logger:        serverLogger,
		mongoDatabase: mongoDatabase,
	}
}
