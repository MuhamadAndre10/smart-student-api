package usecase

import (
	"context"
	"errors"
	database "github.com/MuhamadAndre/auth-service/internal/db"
	"github.com/MuhamadAndre/auth-service/internal/models"
	"github.com/MuhamadAndre/auth-service/internal/models/converter"
	"github.com/MuhamadAndre/auth-service/internal/utils"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"go.uber.org/zap"
	"time"
)

type AuthUseCase struct {
	Log      *zap.Logger
	DB       *database.Queries
	Validate *validator.Validate
}

func NewAuthUseCase(log *zap.Logger, db *database.Queries, validate *validator.Validate) *AuthUseCase {
	return &AuthUseCase{Log: log, DB: db, Validate: validate}
}

func (c *AuthUseCase) SignIn(ctx context.Context, payload *models.LoginRequest) (*models.UserResponse, error) {
	if err := c.Validate.Struct(payload); err != nil {
		c.Log.Error("Error validating request", zap.Error(err))
		return nil, errors.New("invalid request")
	}

	byEmail, err := c.DB.GetUserByEmail(ctx, payload.Email)
	if err != nil {
		c.Log.Error("Error get user by email", zap.Error(err))
		return nil, err
	}

	if !utils.CheckPasswordHash(payload.Password, byEmail.Password) {
		c.Log.Error("Invalid password")
		return nil, errors.New("invalid password & email")
	}

	return converter.RegisterToResponse(byEmail), nil
}

func (c *AuthUseCase) SignUp(ctx context.Context, payload *models.RegisterRequest) (*models.UserResponse, error) {

	timeout, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	if err := c.Validate.Struct(payload); err != nil {
		c.Log.Error("Error validating request", zap.Error(err))
		return nil, err
	}

	password, err := utils.HashPassword(payload.Password)

	params := database.CreateUserParams{
		ID: pgtype.UUID{
			Bytes: uuid.New(),
			Valid: true,
		},
		Username: pgtype.Text{
			String: payload.Username,
			Valid:  true,
		},
		Email:    payload.Email,
		Password: password,
		CreatedAt: pgtype.Timestamp{
			Time:  time.Now().UTC(),
			Valid: true,
		},
		UpdatedAt: pgtype.Timestamp{
			Time:  time.Now().UTC(),
			Valid: true,
		},
	}

	users, err := c.DB.CreateUser(timeout, &params)

	if err != nil {
		c.Log.Error("email already exist", zap.Error(err))
		return nil, err
	}

	return converter.RegisterToResponse(users), nil
}

func (c *AuthUseCase) VerifyUser(ctx context.Context, email string) (*models.UserResponse, error) {
	byEmail, err := c.DB.GetUserByEmail(ctx, email)
	if err != nil {
		c.Log.Error("error get user by email", zap.Error(err))
		return nil, err
	}

	// update user status
	params := database.UpdateUserParams{
		Email:         byEmail.Email,
		VerifiedEmail: true,
		UpdatedAt: pgtype.Timestamp{
			Time:  time.Now().Local().UTC(),
			Valid: true,
		},
	}

	user, err := c.DB.UpdateUser(ctx, &params)
	if err != nil {
		c.Log.Error("error update user", zap.Error(err))
		return nil, err
	}

	return converter.RegisterToResponse(user), nil
}

func (c *AuthUseCase) DeleteUser(ctx context.Context, email string) {
	_, err := c.DB.DeleteUserByEmail(ctx, email)
	if err != nil {
		c.Log.Error("error delete user", zap.Error(err))
		return
	}

	c.Log.Info("user deleted", zap.String("email", email))
}
