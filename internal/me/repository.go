package me

import (
	"context"

	"github.com/tenSunFree/travel-audio-guide-go/internal/db"
)

type Repository struct {
	queries *db.Queries
}

func NewRepository(queries *db.Queries) *Repository {
	return &Repository{queries: queries}
}

func (r *Repository) GetByID(ctx context.Context, id UUID) (db.Profile, error) {
	return r.queries.GetProfileByID(ctx, id.PG())
}

func (r *Repository) Create(ctx context.Context, id UUID, email *string) (db.Profile, error) {
	return r.queries.CreateProfile(ctx, db.CreateProfileParams{
		ID:    id.PG(),
		Email: email,
	})
}

func (r *Repository) Update(ctx context.Context, arg db.UpdateProfileParams) (db.Profile, error) {
	return r.queries.UpdateProfile(ctx, arg)
}
