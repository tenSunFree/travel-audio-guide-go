package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"

	"github.com/tenSunFree/travel-audio-guide-go/internal/auth"
	"github.com/tenSunFree/travel-audio-guide-go/internal/config"
	"github.com/tenSunFree/travel-audio-guide-go/internal/database"
	"github.com/tenSunFree/travel-audio-guide-go/internal/db"
	"github.com/tenSunFree/travel-audio-guide-go/internal/me"
	"github.com/tenSunFree/travel-audio-guide-go/internal/server"
)

func main() {
	_ = godotenv.Load()

	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))

	cfg, err := config.Load()
	if err != nil {
		log.Error("load config failed", "error", err)
		os.Exit(1)
	}

	ctx := context.Background()
	pool, err := database.NewPostgresPool(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Error("connect database failed", "error", err)
		os.Exit(1)
	}
	defer pool.Close()

	queries := db.New(pool)
	meRepo := me.NewRepository(queries)
	meService := me.NewService(meRepo)
	meHandler := me.NewHandler(meService)

	var verifier *auth.JWTVerifier
	if cfg.SupabaseJWKSURL != "" {
		verifier, err = auth.NewJWTVerifierFromJWKS(cfg.SupabaseJWKSURL)
		if err != nil {
			log.Error("init jwt verifier from jwks failed", "error", err)
			os.Exit(1)
		}
		log.Info("jwt verifier ready", "mode", "ES256/JWKS")
	} else {
		verifier = auth.NewJWTVerifier(cfg.SupabaseJWTSecret)
		log.Info("jwt verifier ready", "mode", "HS256/Secret")
	}

	router := server.NewRouter(log, verifier, meHandler)
	httpServer := &http.Server{
		Addr:              cfg.HTTPAddr,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		log.Info("server started", "addr", cfg.HTTPAddr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("server failed", "error", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		log.Error("server shutdown failed", "error", err)
	}
}
