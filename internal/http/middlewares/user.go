package middlewares

import (
	"context"
	"net/http"
)

const userIdHeader = "X-User-Id"
const userIdContextKey = "user_id"

func UserId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		userId := r.Header.Get(userIdHeader)

		ctx = context.WithValue(ctx, userIdContextKey, userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
