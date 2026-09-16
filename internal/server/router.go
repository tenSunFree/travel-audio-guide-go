package server

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/tenSunFree/travel-audio-guide-go/internal/attractions"
	"github.com/tenSunFree/travel-audio-guide-go/internal/auth"
	"github.com/tenSunFree/travel-audio-guide-go/internal/events"
	"github.com/tenSunFree/travel-audio-guide-go/internal/me"
	"github.com/tenSunFree/travel-audio-guide-go/internal/media"
	"github.com/tenSunFree/travel-audio-guide-go/internal/middleware"
	"github.com/tenSunFree/travel-audio-guide-go/internal/miscellaneous"
	"github.com/tenSunFree/travel-audio-guide-go/internal/tours"
	"github.com/tenSunFree/travel-audio-guide-go/pkg/response"
)

func NewRouter(
	log *slog.Logger,
	pool *pgxpool.Pool,
	verifier *auth.JWTVerifier,
	meHandler *me.Handler,
	attractionsHandler *attractions.Handler,
	eventsHandler *events.Handler,
	mediaHandler *media.Handler,
	toursHandler *tours.Handler,
	miscHandler *miscellaneous.Handler,
) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Recovery(log))
	r.Use(middleware.Logger(log))
	r.Use(middleware.CORS)
	r.Use(chimiddleware.RequestID)

	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		response.JSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	r.Get("/readyz", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()
		if err := pool.Ping(ctx); err != nil {
			log.Warn("readiness check failed", "error", err)
			response.Error(w, http.StatusServiceUnavailable, "database not ready")
			return
		}
		response.JSON(w, http.StatusOK, map[string]string{"status": "ready"})
	})

	// Third-party compatible proxy; no Supabase JWT required
	r.Route("/open-api", func(r chi.Router) {
		r.Get("/{lang}/Attractions/All", attractionsHandler.GetAll)
		r.Get("/{lang}/Events/News", eventsHandler.GetNews)
		r.Get("/{lang}/Events/Activity", eventsHandler.GetActivity)
		r.Get("/{lang}/Events/Calendar", eventsHandler.GetCalendar)
		r.Get("/{lang}/Media/Audio", mediaHandler.GetAudio)
		r.Get("/{lang}/Tours/Theme", toursHandler.GetTheme)
		r.Get("/{lang}/Miscellaneous/Categories", miscHandler.GetCategories)
	})

	// Own domain API; Supabase JWT required
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
