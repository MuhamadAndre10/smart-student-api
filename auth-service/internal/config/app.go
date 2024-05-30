package config

import (
	database "github.com/MuhamadAndre/auth-service/internal/db"
	deliverhttp "github.com/MuhamadAndre/auth-service/internal/delivery/http"
	"github.com/MuhamadAndre/auth-service/internal/delivery/http/route"
	"github.com/MuhamadAndre/auth-service/internal/usecase"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
)

type BootstrapConfig struct {
	DB       *database.Queries
	Log      *zap.Logger
	Validate *validator.Validate
}

func Bootstrap(config *BootstrapConfig) http.Handler {

	//repo

	// usecae
	authUseCase := usecase.NewAuthUseCase(config.Log, config.DB, config.Validate)

	// controller
	authController := deliverhttp.NewAuthController(config.Log, authUseCase)

	return route.New(&route.Config{
		AuthController: authController,
	})

}
