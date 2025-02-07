package middlewares

import (
	"context"
	"example/graphql/utils"
	"net/http"
)

func AuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			header := r.Header.Get("Authorization")
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}
			username, err := utils.ParseToken(header)
			if err != nil {
				http.Error(w, "invalid token", http.StatusUnauthorized)
			}
			ctx := context.WithValue(r.Context(), "username", username)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
