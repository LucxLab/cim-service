package http

import (
	cim "github.com/LucxLab/cim-service/internal"
	"github.com/LucxLab/cim-service/internal/http/handler/cdr"
	"github.com/LucxLab/cim-service/internal/http/handler/health"
	"github.com/gorilla/mux"
)

const apiV1Path = "/api/v1"

func setupRouter(uploadService cim.UploadService) *mux.Router {
	router := mux.NewRouter()
	router.StrictSlash(true)

	healthHandler := health.NewHandler(router)
	router.PathPrefix(health.HandlerPath).Handler(healthHandler)

	// Define the router for the API v1
	apiV1 := router.PathPrefix(apiV1Path).Subrouter()

	cdrHandler := cdr.NewHandler(apiV1, uploadService)
	apiV1.PathPrefix(cdr.HandlerPath).Handler(cdrHandler)
	return router
}
