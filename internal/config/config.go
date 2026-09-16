package config

import (
	"fmt"
	"os"
)

type Config struct {
	AppEnv              string
	HTTPAddr            string
	DatabaseURL         string
	SupabaseJWTSecret   string
	SupabaseJWKSURL     string
	TaipeiTravelBaseURL string
}

func Load() (Config, error) {
	cfg := Config{
		AppEnv:              getEnv("APP_ENV", "local"),
		HTTPAddr:            httpAddr(),
		DatabaseURL:         os.Getenv("DATABASE_URL"),
		SupabaseJWTSecret:   os.Getenv("SUPABASE_JWT_SECRET"),
		SupabaseJWKSURL:     os.Getenv("SUPABASE_JWKS_URL"),
		TaipeiTravelBaseURL: getEnv("TAIPEI_TRAVEL_BASE_URL", "https://www.travel.taipei/open-api"),
	}
	if cfg.DatabaseURL == "" {
		return Config{}, fmt.Errorf("DATABASE_URL is required")
	}
	if cfg.SupabaseJWKSURL == "" && cfg.SupabaseJWTSecret == "" {
		return Config{}, fmt.Errorf("SUPABASE_JWKS_URL or SUPABASE_JWT_SECRET is required")
	}
	return cfg, nil
}

// Cloud Run injects PORT; local dev keeps using HTTP_ADDR (or defaults to :8080).
func httpAddr() string {
	if port := os.Getenv("PORT"); port != "" {
		return ":" + port
	}
	if addr := os.Getenv("HTTP_ADDR"); addr != "" {
		return addr
	}
	return ":8080"
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
