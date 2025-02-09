package http

import (
	"github.com/LucxLab/cim-service/internal/cdr"
	"github.com/LucxLab/cim-service/internal/http/handler"
	"github.com/LucxLab/cim-service/internal/http/middleware"
	"net/http"
)

func NewRouter(cdrService cdr.Service) *http.ServeMux {
	router := http.NewServeMux()

	// Handlers
	healthHandler := handler.NewHealthHandler()
	router.HandleFunc("GET /health", healthHandler.GlobalStatus)

	apiV1Router := http.NewServeMux()
	apiV1RouterWithMiddlewares := ChainMiddlewares(apiV1Router, middleware.CorrelationId, middleware.UserId)
	router.Handle("/api/v1/", http.StripPrefix("/api/v1", apiV1RouterWithMiddlewares))

	cdrHandler := handler.NewCdrHandler(cdrService)
	apiV1Router.HandleFunc("POST /organizations/{organization_id}/cdr-files", cdrHandler.UploadCdrFile)
	return router
}

// ChainMiddlewares chains multiple middlewares to a single handler
func ChainMiddlewares(handler http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for _, m := range middlewares {
		handler = m(handler)
	}
	return handler
}
