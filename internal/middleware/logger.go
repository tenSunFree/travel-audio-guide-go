package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

// Logger returns a middleware that logs HTTP requests with method, path, status, and duration.
func Logger(log *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			ww := &wrappedWriter{ResponseWriter: w, status: http.StatusOK}
			next.ServeHTTP(ww, r)
			log.Info("http",
				"method", r.Method,
				"path", r.URL.Path,
				"status", ww.status,
				"duration_ms", time.Since(start).Milliseconds(),
			)
		})
	}
}

type wrappedWriter struct {
	http.ResponseWriter
	status int
}

// WriteHeader captures the status code before delegating to the underlying ResponseWriter.
func (w *wrappedWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}
