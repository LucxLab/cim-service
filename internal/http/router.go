package http

import (
	"github.com/LucxLab/cim-service/internal/cdr"
	"github.com/LucxLab/cim-service/internal/http/handler"
	"net/http"
)

func NewRouter(cdrService cdr.Service) *http.ServeMux {
	router := http.NewServeMux()
	healthHandler := handler.NewHealthHandler()
	router.HandleFunc("GET /health", healthHandler.GlobalStatus)

	apiV1Router := http.NewServeMux()
	router.Handle("/api/v1/", http.StripPrefix("/api/v1", apiV1Router))

	cdrHandler := handler.NewCdrHandler(cdrService)
	apiV1Router.HandleFunc("POST /organizations/{organization_id}/cdr-files", cdrHandler.UploadCdrFile)
	return router
}
