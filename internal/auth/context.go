package auth

import (
	"context"
	"fmt"
)

type contextKey string

const userContextKey contextKey = "auth_user"

type User struct {
	ID    string
	Email string
	Role  string
}

// WithUser adds an authenticated user to the context.
func WithUser(ctx context.Context, user User) context.Context {
	return context.WithValue(ctx, userContextKey, user)
}

// UserFromContext retrieves the authenticated user from the context.
func UserFromContext(ctx context.Context) (User, error) {
	user, ok := ctx.Value(userContextKey).(User)
	if !ok || user.ID == "" {
		return User{}, fmt.Errorf("auth user not found in context")
	}
	return user, nil
}
