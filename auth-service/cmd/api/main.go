package main

import (
	"fmt"
	"github.com/MuhamadAndre/auth-service/internal/config"
	"github.com/MuhamadAndre/auth-service/internal/utils"
	"go.uber.org/zap"
	"net/http"
)

func main() {
	utils.InitEnvConfigs(".")

	log := config.NewLogger()
	db := config.OpenConnectionDB(log)
	validate := config.NewValidate()

	cfg := &config.BootstrapConfig{
		DB:       db,
		Log:      log,
		Validate: validate,
	}

	srv := &http.Server{
		Handler: config.Bootstrap(cfg),
		Addr:    fmt.Sprintf(":%s", utils.Env.AppPort),
	}

	log.Info("Server started", zap.String("port", utils.Env.AppPort))

	err := srv.ListenAndServe()
	if err != nil {
		log.Error("Error starting server", zap.Error(err))
		panic(err)
	}

}
