package route

import (
	"github.com/MuhamadAndre10/student-profile-service/internal/delivery/http"
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	App            *fiber.App
	StudentHandler *http.StudentHandler
}

func (r *Config) Setup() {
	r.App.Post("/student", r.StudentHandler.Insert)
}
