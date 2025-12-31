package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/FelipePn10/panossoerp/internal/infrastructure/auth"
	contextkey "github.com/FelipePn10/panossoerp/internal/interfaces/http/context"
	"github.com/golang-jwt/jwt/v5"
)

type Logger interface {
	Warn(msg string, args ...any)
}

func JWT(secret string, logger Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodOptions {
				next.ServeHTTP(w, r)
				return
			}

			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, `{"error": "Authorization header missing"}`, http.StatusUnauthorized)
				return
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, `{"error": "Invalid auth header format"}`, http.StatusUnauthorized)
				return
			}

			claims := &auth.UserClaims{}

			token, err := jwt.ParseWithClaims(
				parts[1],
				claims,
				func(t *jwt.Token) (interface{}, error) {
					if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, jwt.ErrSignatureInvalid
					}
					return []byte(secret), nil
				},
			)

			if err != nil || !token.Valid {
				logger.Warn("invalid token attempt", "error", err)
				http.Error(w, `{"error": "Invalid token"}`, http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(
				r.Context(),
				contextkey.UserKey,
				claims,
			)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func RequireRole(roles ...string) func(http.Handler) http.Handler {
	roleSet := make(map[string]struct{})
	for _, r := range roles {
		roleSet[r] = struct{}{}
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims, ok := r.Context().Value(contextkey.UserKey).(*auth.UserClaims)
			if !ok {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}

			if _, allowed := roleSet[claims.Role]; !allowed {
				http.Error(w, "forbidden", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
