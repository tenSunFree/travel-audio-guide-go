package me

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/jackc/pgx/v5"

	"github.com/tenSunFree/travel-audio-guide-go/internal/auth"
	"github.com/tenSunFree/travel-audio-guide-go/internal/db"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

type NotFoundError struct{ Message string }

func (e *NotFoundError) Error() string { return e.Message }

func (s *Service) GetMe(ctx context.Context, user auth.User) (ProfileResponse, error) {
	userID, err := ParseUUID(user.ID)
	if err != nil {
		return ProfileResponse{}, err
	}

	profile, err := s.repo.GetByID(ctx, userID)
	if err == nil {
		return toResponse(profile), nil
	}
	if !errors.Is(err, pgx.ErrNoRows) {
		return ProfileResponse{}, fmt.Errorf("get profile: %w", err)
	}

	email := optionalString(user.Email)
	profile, err = s.repo.Create(ctx, userID, email)
	if err != nil {
		return ProfileResponse{}, fmt.Errorf("create profile: %w", err)
	}
	return toResponse(profile), nil
}

func (s *Service) UpdateMe(ctx context.Context, user auth.User, req UpdateMeRequest) (ProfileResponse, error) {
	userID, err := ParseUUID(user.ID)
	if err != nil {
		return ProfileResponse{}, err
	}

	if err := validateUpdateRequest(&req); err != nil {
		return ProfileResponse{}, err
	}

	profile, err := s.repo.Update(ctx, db.UpdateProfileParams{
		ID:                userID.PG(),
		DisplayName:       req.DisplayName,
		AvatarUrl:         req.AvatarURL,
		PreferredLanguage: req.PreferredLanguage,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ProfileResponse{}, &NotFoundError{Message: "profile not found, call GET /api/v1/me first"}
		}
		return ProfileResponse{}, fmt.Errorf("update profile: %w", err)
	}

	return toResponse(profile), nil
}

var allowedLanguages = map[string]bool{
	"zh-TW": true,
	"zh-CN": true,
	"en":    true,
	"ja":    true,
}

func validateUpdateRequest(req *UpdateMeRequest) error {
	if req.DisplayName != nil {
		trimmed := strings.TrimSpace(*req.DisplayName)
		if len([]rune(trimmed)) < 1 || len([]rune(trimmed)) > 50 {
			return &ValidationError{Field: "display_name", Message: "must be 1-50 characters"}
		}
		req.DisplayName = &trimmed
	}

	if req.AvatarURL != nil && strings.TrimSpace(*req.AvatarURL) != "" {
		parsed, err := url.ParseRequestURI(*req.AvatarURL)
		if err != nil || parsed.Scheme == "" || parsed.Host == "" {
			return &ValidationError{Field: "avatar_url", Message: "must be a valid URL (https://...)"}
		}
	}

	if req.PreferredLanguage != nil {
		if !allowedLanguages[*req.PreferredLanguage] {
			return &ValidationError{
				Field:   "preferred_language",
				Message: "must be one of: zh-TW, zh-CN, en, ja",
			}
		}
	}

	return nil
}

func optionalString(v string) *string {
	if v == "" {
		return nil
	}
	return &v
}

func toResponse(p db.Profile) ProfileResponse {
	return ProfileResponse{
		ID:                uuidToString(p.ID),
		Email:             p.Email,
		DisplayName:       p.DisplayName,
		AvatarURL:         p.AvatarUrl,
		PreferredLanguage: p.PreferredLanguage,
		CreatedAt:         p.CreatedAt,
		UpdatedAt:         p.UpdatedAt,
	}
}
