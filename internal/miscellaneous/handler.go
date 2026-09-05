package miscellaneous

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

// NewHandler creates a new Handler for miscellaneous endpoints.
func NewHandler(service *Service, log *slog.Logger) *Handler {
	return &Handler{service: service, log: log}
}

// GetCategories handles GET /{lang}/Miscellaneous/Categories requests to retrieve categories.
func (h *Handler) GetCategories(w http.ResponseWriter, r *http.Request) {
	lang := chi.URLParam(r, "lang")
	typeParam := r.URL.Query().Get("type")

	result, err := h.service.GetCategories(r.Context(), lang, typeParam, r.URL.RawQuery)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if err != nil {
		var valErr *ValidationError
		if errors.As(err, &valErr) {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": valErr.Error()})
			return
		}

		h.log.Error("get miscellaneous categories failed", "lang", lang, "type", typeParam, "query", r.URL.RawQuery, "error", err)
		w.WriteHeader(http.StatusBadGateway)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": "upstream service unavailable"})
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		h.log.Error("encode miscellaneous categories response failed", "error", err)
	}
}
