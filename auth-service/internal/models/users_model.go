package models

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type UserResponse struct {
	ID            pgtype.UUID      `json:"id"`
	FullName      pgtype.Text      `json:"full_name"`
	Username      pgtype.Text      `json:"username"`
	Email         string           `json:"email"`
	Photo         pgtype.Text      `json:"photo"`
	UserActive    bool             `json:"user_active"`
	VerifiedEmail bool             `json:"verified_email"`
	Password      string           `json:"password"`
	CreatedAt     pgtype.Timestamp `json:"created_at"`
	UpdatedAt     pgtype.Timestamp `json:"updated_at"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterResponse struct {
	ID            pgtype.UUID      `json:"id"`
	Username      pgtype.Text      `json:"username"`
	Email         string           `json:"email"`
	Token         string           `json:"token"`
	VerifiedEmail bool             `json:"verified_email"`
	CreatedAt     pgtype.Timestamp `json:"created_at"`
	UpdatedAt     pgtype.Timestamp `json:"updated_at"`
}
