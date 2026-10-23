package config

import "testing"

func TestHTTPAddr(t *testing.T) {
	tests := []struct {
		name     string
		port     string
		httpAddr string
		want     string
	}{
		{name: "default", want: ":8080"},
		{name: "HTTP_ADDR only", httpAddr: ":9000", want: ":9000"},
		{name: "PORT only", port: "3000", want: ":3000"},
		{name: "PORT wins over HTTP_ADDR", port: "3000", httpAddr: ":9000", want: ":3000"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("PORT", tt.port)
			t.Setenv("HTTP_ADDR", tt.httpAddr)

			if got := httpAddr(); got != tt.want {
				t.Fatalf("httpAddr() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestLoad(t *testing.T) {
	tests := []struct {
		name        string
		databaseURL string
		jwksURL     string
		jwtSecret   string
		wantErr     bool
	}{
		{name: "missing database url", jwtSecret: "secret", wantErr: true},
		{name: "missing supabase verifier", databaseURL: "postgres://x", wantErr: true},
		{name: "jwks only", databaseURL: "postgres://x", jwksURL: "https://example.com/jwks.json"},
		{name: "secret only", databaseURL: "postgres://x", jwtSecret: "secret"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("DATABASE_URL", tt.databaseURL)
			t.Setenv("SUPABASE_JWKS_URL", tt.jwksURL)
			t.Setenv("SUPABASE_JWT_SECRET", tt.jwtSecret)

			_, err := Load()
			if (err != nil) != tt.wantErr {
				t.Fatalf("Load() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLoadDefaults(t *testing.T) {
	t.Setenv("DATABASE_URL", "postgres://x")
	t.Setenv("SUPABASE_JWT_SECRET", "secret")
	t.Setenv("SUPABASE_JWKS_URL", "")
	t.Setenv("APP_ENV", "")
	t.Setenv("TAIPEI_TRAVEL_BASE_URL", "")
	t.Setenv("PORT", "")
	t.Setenv("HTTP_ADDR", "")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	if cfg.AppEnv != "local" {
		t.Errorf("AppEnv = %q, want %q", cfg.AppEnv, "local")
	}
	if cfg.HTTPAddr != ":8080" {
		t.Errorf("HTTPAddr = %q, want %q", cfg.HTTPAddr, ":8080")
	}
	if cfg.TaipeiTravelBaseURL != "https://www.travel.taipei/open-api" {
		t.Errorf("TaipeiTravelBaseURL = %q", cfg.TaipeiTravelBaseURL)
	}
}
