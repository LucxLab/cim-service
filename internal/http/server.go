package http

import (
	"fmt"
	cim "github.com/LucxLab/cim-service/internal"
	upload "github.com/LucxLab/cim-service/internal/cdr"
	"github.com/LucxLab/cim-service/internal/mongo"
	"github.com/LucxLab/cim-service/internal/mongo/cdr"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	Address  string
	database cim.Database
	router   *mux.Router
}

func (s *Server) Listen() error {
	fmt.Println("Server listening on", s.Address)
	return http.ListenAndServe(s.Address, s.router)
}

func (s *Server) Close() error {
	return s.database.Close()
}

func NewServer(address string) *Server {
	database := mongo.NewDatabase()

	// Repositories
	cdrRepository := cdr.NewRepository(database)

	// Services
	uploadService := upload.NewService(cdrRepository)

	router := setupRouter(uploadService)
	return &Server{
		Address:  address,
		database: database,
		router:   router,
	}
}
