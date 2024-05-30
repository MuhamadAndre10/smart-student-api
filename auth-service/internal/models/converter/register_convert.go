package converter

import (
	database "github.com/MuhamadAndre/auth-service/internal/db"
	"github.com/MuhamadAndre/auth-service/internal/models"
)

func RegisterToResponse(user *database.User) *models.UserResponse {
	return &models.UserResponse{
		ID:            user.ID,
		Email:         user.Email,
		Username:      user.Username,
		Photo:         user.Photo,
		FullName:      user.FullName,
		VerifiedEmail: user.VerifiedEmail,
		Password:      user.Password,
		UpdatedAt:     user.UpdatedAt,
		UserActive:    user.UserActive,
		CreatedAt:     user.CreatedAt,
	}
}
