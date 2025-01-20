package handler

import (
	"github.com/LucxLab/cim-service/internal/health"
	"net/http"
)

type handler struct{}

func (h handler) GlobalStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"status": "ok"}`))
}

func NewHealthHandler() health.Handler {
	return &handler{}
}
