package models

import "placid-backend/internal/database"

type RegisterRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserRequest struct {
	Email string `json:"email" validate:"required"`
}

type DeleteUserRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type AdminUploadTrackRequest struct {
	Email string         `json:"email" validate:"required"`
	Track database.Track `json:"track" validate:"required"`
}

type AdminDeleteTrackRequest struct {
	Email string         `json:"email" validate:"required"`
	Track database.Track `json:"track" validate:"required"`
}
type SubscribeToNewsletterRequest struct {
	Email string `json:"email" validate:"required"`
}
type UnsubscribeFromNewsletterRequest struct {
	Email string `json:"email" validate:"required"`
}
