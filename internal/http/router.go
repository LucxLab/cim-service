package http

import (
	"github.com/LucxLab/cim-service/internal/http/handlers"
	"github.com/LucxLab/cim-service/internal/http/middlewares"
	"github.com/LucxLab/cim-service/internal/upload"
	"net/http"
)

func NewRouter(uploadUseCase upload.UseCase) *http.ServeMux {
	router := http.NewServeMux()

	// Handlers
	healthHandler := handlers.NewHealth()
	router.HandleFunc("GET /health", healthHandler.GlobalStatus)

	apiV1Router := http.NewServeMux()
	apiV1RouterWithMiddlewares := ChainMiddlewares(apiV1Router, middlewares.CorrelationId, middlewares.UserId)
	router.Handle("/api/v1/", http.StripPrefix("/api/v1", apiV1RouterWithMiddlewares))

	cdrHandler := handlers.NewCdr(uploadUseCase)
	apiV1Router.HandleFunc("POST /organizations/{organization_id}/cdr-files", cdrHandler.UploadFile)
	return router
}

// ChainMiddlewares chains multiple middlewares to a single handler
func ChainMiddlewares(handler http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for _, m := range middlewares {
		handler = m(handler)
	}
	return handler
}
