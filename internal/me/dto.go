package me

import "time"

type ProfileResponse struct {
	ID                string    `json:"id"`
	Email             *string   `json:"email"`
	DisplayName       *string   `json:"display_name"`
	AvatarURL         *string   `json:"avatar_url"`
	PreferredLanguage string    `json:"preferred_language"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type UpdateMeRequest struct {
	DisplayName       *string `json:"display_name"`
	AvatarURL         *string `json:"avatar_url"`
	PreferredLanguage *string `json:"preferred_language"`
}
