package config

import (
	"fmt"
	"os"
)

type Config struct {
	AppEnv            string
	HTTPAddr          string
	DatabaseURL       string
	SupabaseJWTSecret string
	SupabaseJWKSURL   string
}

func Load() (Config, error) {
	cfg := Config{
		AppEnv:            getEnv("APP_ENV", "local"),
		HTTPAddr:          getEnv("HTTP_ADDR", ":8080"),
		DatabaseURL:       os.Getenv("DATABASE_URL"),
		SupabaseJWTSecret: os.Getenv("SUPABASE_JWT_SECRET"),
		SupabaseJWKSURL:   os.Getenv("SUPABASE_JWKS_URL"),
	}
	if cfg.DatabaseURL == "" {
		return Config{}, fmt.Errorf("DATABASE_URL is required")
	}
	if cfg.SupabaseJWKSURL == "" && cfg.SupabaseJWTSecret == "" {
		return Config{}, fmt.Errorf("SUPABASE_JWKS_URL or SUPABASE_JWT_SECRET is required")
	}
	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
