package middleware

import (
	"context"
	"github.com/LucxLab/cim-service/internal/constant"
	"net/http"
)

const userIdHeader = "X-User-Id"

func UserId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		userId := r.Header.Get(userIdHeader)

		ctx = context.WithValue(ctx, constant.UserIdContextKey, userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
