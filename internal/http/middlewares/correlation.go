package middlewares

import (
	"context"
	"github.com/google/uuid"
	"net/http"
)

const correlationIdHeader = "X-Correlation-Id"
const correlationIdContextKey = "correlation_id"

func CorrelationId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		correlationId := r.Header.Get(correlationIdHeader)
		if correlationId == "" {
			correlationId = uuid.New().String()
		}

		ctx = context.WithValue(ctx, correlationIdContextKey, correlationId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
