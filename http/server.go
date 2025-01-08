package http

import (
	"fmt"
	"github.com/LucxLab/cim-service/http/handler/health"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	Address string
}

func (s *Server) Listen() error {
	router := mux.NewRouter()
	router.StrictSlash(true)

	healthHandler := health.NewHandler(router)
	router.PathPrefix(health.HandlerPath).Handler(healthHandler)

	fmt.Println("Server listening on", s.Address)
	return http.ListenAndServe(s.Address, router)
}
