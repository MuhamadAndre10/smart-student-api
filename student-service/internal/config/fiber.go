package config

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func NewFiber(viper *viper.Viper) *fiber.App {
	return fiber.New(fiber.Config{
		AppName:      viper.GetString("APP_NAME"),
		ErrorHandler: NewErrorHandler(),
	})
}

func NewErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		var e *fiber.Error
		if errors.As(err, &e) {
			code = e.Code
		}

		return ctx.Status(code).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}
}
