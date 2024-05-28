package config

import (
	deliver_http "github.com/MuhamadAndre/auth-service/internal/delivery/http"
	"github.com/MuhamadAndre/auth-service/internal/delivery/http/route"
	"github.com/MuhamadAndre/auth-service/internal/usecase"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
	"net/http"
)

type BootstrapConfig struct {
	DB  *pgx.Conn
	Log *zap.Logger
}

func Bootstrap(config *BootstrapConfig) http.Handler {

	//repo

	// usecae
	authUseCase := usecase.NewAuthUseCase(config.Log)

	// controller
	authController := deliver_http.NewAuthController(config.Log, authUseCase)

	return route.New(&route.Config{
		AuthController: authController,
	})

}
