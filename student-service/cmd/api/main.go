package main

import (
	"fmt"
	"github.com/MuhamadAndre10/student-profile-service/internal/config"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config.InitEnvConfigs(".")
	app := config.NewFiber()
	log := config.NewLogger()
	db := config.NewDatabase()
	validate := config.NewValidate()

	config.Bootstrap(&config.BootstrapConfig{
		Log:      log,
		DB:       db,
		Validate: validate,
		App:      app,
	})

	serve(app, log)
}

func serve(app *fiber.App, log *zap.Logger) {

	go func() {
		err := app.Listen(fmt.Sprintf(":%s", config.Env.AppPort))
		if err != nil {
			log.Panic("Error when listen", zap.Error(err))
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	_ = <-c // block until receive signal
	log.Info("Shutting down the server...")
	err := app.Shutdown()
	if err != nil {
		log.Error("Error when shutdown", zap.Error(err))
	}

	log.Info("Server is shut down")

}
