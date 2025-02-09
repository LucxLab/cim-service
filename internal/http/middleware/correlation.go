package middleware

import (
	"context"
	"github.com/LucxLab/cim-service/internal/constant"
	"github.com/google/uuid"
	"net/http"
)

const correlationIdHeader = "X-Correlation-Id"

func CorrelationId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		correlationId := r.Header.Get(correlationIdHeader)
		if correlationId == "" {
			correlationId = uuid.New().String()
		}

		ctx = context.WithValue(ctx, constant.CorrelationIdContextKey, correlationId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
