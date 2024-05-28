package config

import (
	"github.com/MuhamadAndre10/student-profile-service/internal/delivery/http"
	"github.com/MuhamadAndre10/student-profile-service/internal/delivery/http/route"
	"github.com/MuhamadAndre10/student-profile-service/internal/repository"
	"github.com/MuhamadAndre10/student-profile-service/internal/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	App      *fiber.App
	Log      *zap.Logger
	DB       *gorm.DB
	Validate *validator.Validate
}

func Bootstrap(config *BootstrapConfig) {

	// repository
	studentRepo := repository.NewStudentRepository()

	// useCase
	studentUseCase := usecase.NewStudentUseCase(config.DB, config.Log, config.Validate, studentRepo)

	// handler
	studentHandler := http.NewStudentHandler(config.Log, studentUseCase)

	r := route.Config{
		App:            config.App,
		StudentHandler: studentHandler,
	}

	r.Setup()
}
