package auth

import (
	"context"
	"dating-app/internal/helpers"
	"net/http"
	"strings"
)

type contextKey string

const EmailKey contextKey = "email"

// Middleware: to validate access
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		givenToken := r.Header.Get("Authorization")
		givenToken = strings.TrimPrefix(givenToken, "Bearer ")
		claims, err := validateToken(givenToken)
		if err != nil || claims == nil {
			helpers.WriteJSONResponse(w, http.StatusUnauthorized, err)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, EmailKey, claims.Email)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
