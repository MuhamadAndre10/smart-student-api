package usecase

import "go.uber.org/zap"

type AuthUseCase struct {
	Log *zap.Logger
}

func NewAuthUseCase(log *zap.Logger) *AuthUseCase {
	return &AuthUseCase{Log: log}
}

func (c *AuthUseCase) SignIn() {
	c.Log.Info("Sign In")
}
