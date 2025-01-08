package health

import (
	"github.com/gorilla/mux"
	"net/http"
)

const HandlerPath = "/health"

type Handler struct {
	*mux.Router
}

func NewHandler(router *mux.Router) *Handler {
	healthRouter := router.PathPrefix(HandlerPath).Subrouter()
	handler := &Handler{
		Router: healthRouter,
	}

	healthRouter.HandleFunc("", handler.globalStatus).Methods(http.MethodGet)
	return handler
}

func (h *Handler) globalStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"status": "ok"}`))
}
