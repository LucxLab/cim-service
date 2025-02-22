package handlers

import (
	"github.com/LucxLab/cim-service/internal/health"
	"net/http"
)

type handler struct{}

func (h handler) GlobalStatus(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"status": "ok"}`))
}

func NewHealth() health.Handler {
	return &handler{}
}
