package server

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

// newTestRouter builds the real router with nil dependencies. That is enough for
// routes that never touch the database, the JWT verifier, or the upstream API.
func newTestRouter() http.Handler {
	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	return NewRouter(log, nil, nil, nil, nil, nil, nil, nil, nil)
}

func doRequest(t *testing.T, h http.Handler, method, path string, headers map[string]string) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequest(method, path, nil)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	return rec
}

func TestHealth(t *testing.T) {
	rec := doRequest(t, newTestRouter(), http.MethodGet, "/health", nil)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	var body struct {
		Success bool `json:"success"`
		Data    struct {
			Status string `json:"status"`
		} `json:"data"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("invalid JSON body: %v", err)
	}
	if !body.Success || body.Data.Status != "ok" {
		t.Fatalf("body = %s, want success=true and data.status=ok", rec.Body.String())
	}
}

// Cloud Run reserves some paths ending in "z", so the old names must stay unregistered.
func TestLegacyHealthPathsAreNotRegistered(t *testing.T) {
	for _, path := range []string{"/healthz", "/readyz"} {
		rec := doRequest(t, newTestRouter(), http.MethodGet, path, nil)
		if rec.Code != http.StatusNotFound {
			t.Errorf("GET %s status = %d, want %d", path, rec.Code, http.StatusNotFound)
		}
	}
}

func TestProtectedRouteRequiresBearerToken(t *testing.T) {
	tests := []struct {
		name    string
		headers map[string]string
		wantErr string
	}{
		{name: "no header", wantErr: "missing authorization header"},
		{
			name:    "wrong scheme",
			headers: map[string]string{"Authorization": "Basic abc"},
			wantErr: "authorization must be Bearer token",
		},
		{
			name:    "empty token",
			headers: map[string]string{"Authorization": "Bearer  "},
			wantErr: "empty bearer token",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := doRequest(t, newTestRouter(), http.MethodGet, "/api/v1/me", tt.headers)

			if rec.Code != http.StatusUnauthorized {
				t.Fatalf("status = %d, want %d", rec.Code, http.StatusUnauthorized)
			}

			var body struct {
				Success bool   `json:"success"`
				Error   string `json:"error"`
			}
			if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
				t.Fatalf("invalid JSON body: %v", err)
			}
			if body.Success || body.Error != tt.wantErr {
				t.Fatalf("body = %s, want error %q", rec.Body.String(), tt.wantErr)
			}
		})
	}
}

func TestUnknownRouteReturnsJSON404(t *testing.T) {
	rec := doRequest(t, newTestRouter(), http.MethodGet, "/does-not-exist", nil)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusNotFound)
	}
	if ct := rec.Header().Get("Content-Type"); ct != "application/json; charset=utf-8" {
		t.Fatalf("Content-Type = %q", ct)
	}
}
