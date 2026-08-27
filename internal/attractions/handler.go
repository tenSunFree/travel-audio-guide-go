package attractions

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

func NewHandler(service *Service, log *slog.Logger) *Handler {
	return &Handler{service: service, log: log}
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	lang := chi.URLParam(r, "lang")

	result, err := h.service.GetAll(r.Context(), lang, r.URL.RawQuery)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		var valErr *ValidationError
		if errors.As(err, &valErr) {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": valErr.Error()})
			return
		}

		h.log.Error("get attractions failed", "lang", lang, "query", r.URL.RawQuery, "error", err)
		w.WriteHeader(http.StatusBadGateway)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": "upstream service unavailable"})
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		h.log.Error("encode attractions response failed", "error", err)
	}
}
