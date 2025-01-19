package http

import (
	"github.com/LucxLab/cim-service/internal/cdr"
	"github.com/LucxLab/cim-service/internal/http/handler"
	"net/http"
)

func NewRouter(cdrService cdr.Service) *http.ServeMux {
	router := http.NewServeMux()

	cdrHandler := handler.NewCdrHandler(cdrService)
	router.HandleFunc("POST /cdr", cdrHandler.UploadFile)

	return router
}
