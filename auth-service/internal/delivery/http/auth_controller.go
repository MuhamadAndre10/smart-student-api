package deliver_http

import (
	"github.com/MuhamadAndre/auth-service/internal/usecase"
	"go.uber.org/zap"
	"net/http"
)

type AuthController struct {
	log         *zap.Logger
	authUseCase *usecase.AuthUseCase
}

func NewAuthController(log *zap.Logger, authUseCase *usecase.AuthUseCase) *AuthController {
	return &AuthController{
		log:         log,
		authUseCase: authUseCase,
	}
}

func (c *AuthController) SignIn(w http.ResponseWriter, r *http.Request) {
	c.authUseCase.SignIn()
}
