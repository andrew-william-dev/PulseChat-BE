package middleware

import (
	"context"
	"net/http"
	"strings"

	"chatapp/utils"
)

func SocketMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		var token string

		if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
			token = strings.TrimPrefix(authHeader, "Bearer ")
		} else {
			// Fallback: Read from query param for WebSocket
			token = r.URL.Query().Get("token")
		}

		if token == "" {
			http.Error(w, "Missing auth token", http.StatusUnauthorized)
			return
		}

		userID, err := utils.VerifyJWT(token)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
