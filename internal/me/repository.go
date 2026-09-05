package me

import (
	"context"

	"github.com/tenSunFree/travel-audio-guide-go/internal/db"
)

type Repository struct {
	queries *db.Queries
}

// NewRepository creates a new Repository for user profile database operations.
func NewRepository(queries *db.Queries) *Repository {
	return &Repository{queries: queries}
}

// GetByID retrieves a profile by user ID.
func (r *Repository) GetByID(ctx context.Context, id UUID) (db.Profile, error) {
	return r.queries.GetProfileByID(ctx, id.PG())
}

// Create creates a new profile with the given user ID and email.
func (r *Repository) Create(ctx context.Context, id UUID, email *string) (db.Profile, error) {
	return r.queries.CreateProfile(ctx, db.CreateProfileParams{
		ID:    id.PG(),
		Email: email,
	})
}

// Update updates an existing profile with the provided parameters.
func (r *Repository) Update(ctx context.Context, arg db.UpdateProfileParams) (db.Profile, error) {
	return r.queries.UpdateProfile(ctx, arg)
}
