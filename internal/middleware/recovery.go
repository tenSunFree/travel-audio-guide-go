package middleware

import (
	"log/slog"
	"net/http"
	"runtime/debug"

	"github.com/tenSunFree/travel-audio-guide-go/pkg/response"
)

// Recovery returns a middleware that recovers from panics and logs the error with stack trace.
func Recovery(log *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rec := recover(); rec != nil {
					log.Error("panic recovered",
						"error", rec,
						"path", r.URL.Path,
						"stack", string(debug.Stack()),
					)
					response.Error(w, http.StatusInternalServerError, "internal server error")
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
