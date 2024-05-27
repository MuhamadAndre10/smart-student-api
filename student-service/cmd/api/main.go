package main

import (
	"fmt"
	"github.com/MuhamadAndre10/student-profile-service/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.NewViper()
	app := config.NewFiber(cfg)
	log := config.NewLogger()
	db := config.NewDatabase(cfg, log)

	serve(app, cfg, log)
}

func serve(app *fiber.App, cfg *viper.Viper, log *zap.Logger) {

	go func() {
		err := app.Listen(fmt.Sprintf(":%s", cfg.GetString("PORT")))
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
