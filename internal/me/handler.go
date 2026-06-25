package me

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/tenSunFree/travel-audio-guide-go/internal/auth"
	"github.com/tenSunFree/travel-audio-guide-go/pkg/response"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetMe(w http.ResponseWriter, r *http.Request) {
	user, err := auth.UserFromContext(r.Context())
	if err != nil {
		response.Error(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	profile, err := h.service.GetMe(r.Context(), user)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to get profile")
		return
	}

	response.JSON(w, http.StatusOK, profile)
}

func (h *Handler) UpdateMe(w http.ResponseWriter, r *http.Request) {
	user, err := auth.UserFromContext(r.Context())
	if err != nil {
		response.Error(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	defer r.Body.Close()
	var req UpdateMeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	profile, err := h.service.UpdateMe(r.Context(), user, req)
	if err != nil {
		var valErr *ValidationError
		if errors.As(err, &valErr) {
			response.Error(w, http.StatusUnprocessableEntity, valErr.Error())
			return
		}
		var notFoundErr *NotFoundError
		if errors.As(err, &notFoundErr) {
			response.Error(w, http.StatusNotFound, notFoundErr.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to update profile")
		return
	}

	response.JSON(w, http.StatusOK, profile)
}
