package middleware

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/tenSunFree/travel-audio-guide-go/internal/auth"
	"github.com/tenSunFree/travel-audio-guide-go/pkg/response"
)

func Auth(verifier *auth.JWTVerifier) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authorization := r.Header.Get("Authorization")
			if authorization == "" {
				response.Error(w, http.StatusUnauthorized, "missing authorization header")
				return
			}

			if !strings.HasPrefix(strings.ToLower(authorization), "bearer ") {
				response.Error(w, http.StatusUnauthorized, "authorization must be Bearer token")
				return
			}
			tokenString := strings.TrimSpace(authorization[7:])
			if tokenString == "" {
				response.Error(w, http.StatusUnauthorized, "empty bearer token")
				return
			}

			claims, err := verifier.Verify(tokenString)
			if err != nil {
				slog.Warn("jwt verification failed", "error", err, "path", r.URL.Path)
				response.Error(w, http.StatusUnauthorized, "invalid or expired token")
				return
			}

			if claims.Role != "authenticated" {
				response.Error(w, http.StatusForbidden, "token role is not authenticated")
				return
			}

			user := auth.User{
				ID:    claims.Subject,
				Email: claims.Email,
				Role:  claims.Role,
			}
			next.ServeHTTP(w, r.WithContext(auth.WithUser(r.Context(), user)))
		})
	}
}
