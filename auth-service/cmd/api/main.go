package main

import (
	"fmt"
	"github.com/MuhamadAndre/auth-service/internal/config"
	"go.uber.org/zap"
	"net/http"
)

func main() {
	config.InitEnvConfigs(".")

	log := config.NewLogger()
	db := config.OpenConnectionDB(log)

	cfg := &config.BootstrapConfig{
		DB:  db,
		Log: log,
	}

	srv := &http.Server{
		Handler: config.Bootstrap(cfg),
		Addr:    fmt.Sprintf(":%s", config.Env.AppPort),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Error("Error starting server", zap.Error(err))
		panic(err)
	}

}
