package server

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"

	"github.com/tenSunFree/travel-audio-guide-go/internal/auth"
	"github.com/tenSunFree/travel-audio-guide-go/internal/me"
	"github.com/tenSunFree/travel-audio-guide-go/internal/middleware"
	"github.com/tenSunFree/travel-audio-guide-go/pkg/response"
)

func NewRouter(log *slog.Logger, verifier *auth.JWTVerifier, meHandler *me.Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Recovery(log))
	r.Use(middleware.Logger(log))
	r.Use(middleware.CORS)
	r.Use(chimiddleware.RequestID)

	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		response.JSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	r.Route("/api/v1", func(r chi.Router) {
		r.Use(middleware.Auth(verifier))
		r.Get("/me", meHandler.GetMe)
		r.Put("/me", meHandler.UpdateMe)
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		response.Error(w, http.StatusNotFound, "route not found")
	})

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		response.Error(w, http.StatusMethodNotAllowed, "method not allowed")
	})

	return r
}
