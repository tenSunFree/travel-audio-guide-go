package response

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type successBody struct {
	Success bool `json:"success"`
	Data    any  `json:"data"`
}

type errorBody struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

func JSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(successBody{Success: true, Data: data}); err != nil {
		slog.Error("encode response failed", "error", err)
	}
}

func Error(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(errorBody{Success: false, Error: message}); err != nil {
		slog.Error("encode error response failed", "error", err)
	}
}
