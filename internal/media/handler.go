package media

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

func (h *Handler) GetAudio(w http.ResponseWriter, r *http.Request) {
	lang := chi.URLParam(r, "lang")
	result, err := h.service.GetAudio(r.Context(), lang, r.URL.RawQuery)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if err != nil {
		var valErr *ValidationError
		if errors.As(err, &valErr) {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": valErr.Error()})
			return
		}

		h.log.Error("get media audio failed", "lang", lang, "query", r.URL.RawQuery, "error", err)
		w.WriteHeader(http.StatusBadGateway)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": "upstream service unavailable"})
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		h.log.Error("encode media audio response failed", "error", err)
	}
}
