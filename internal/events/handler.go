package events

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service *Service
	log     *slog.Logger
}

// NewHandler creates a new Handler for events endpoints.
func NewHandler(service *Service, log *slog.Logger) *Handler {
	return &Handler{service: service, log: log}
}

// GetNews handles GET /{lang}/Events/News requests to retrieve event news.
func (h *Handler) GetNews(w http.ResponseWriter, r *http.Request) {
	lang := chi.URLParam(r, "lang")
	result, err := h.service.GetNews(r.Context(), lang, r.URL.RawQuery)
	h.writeResult(w, "events_news", lang, r.URL.RawQuery, result, err)
}

// GetActivity handles GET /{lang}/Events/Activity requests to retrieve event activities.
func (h *Handler) GetActivity(w http.ResponseWriter, r *http.Request) {
	lang := chi.URLParam(r, "lang")
	result, err := h.service.GetActivity(r.Context(), lang, r.URL.RawQuery)
	h.writeResult(w, "events_activity", lang, r.URL.RawQuery, result, err)
}

// GetCalendar handles GET /{lang}/Events/Calendar requests to retrieve event calendar.
func (h *Handler) GetCalendar(w http.ResponseWriter, r *http.Request) {
	lang := chi.URLParam(r, "lang")
	result, err := h.service.GetCalendar(r.Context(), lang, r.URL.RawQuery)
	h.writeResult(w, "events_calendar", lang, r.URL.RawQuery, result, err)
}

func (h *Handler) writeResult(w http.ResponseWriter, resource string, lang string, query string, result any, err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if err != nil {
		var valErr *ValidationError
		if errors.As(err, &valErr) {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": valErr.Error()})
			return
		}

		h.log.Error("get taipei travel events failed", "resource", resource, "lang", lang, "query", query, "error", err)
		w.WriteHeader(http.StatusBadGateway)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": "upstream service unavailable"})
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		h.log.Error("encode events response failed", "resource", resource, "error", err)
	}
}
