package route

import (
	"github.com/MuhamadAndre10/student-profile-service/internal/delivery/http"
	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App            *fiber.App
	StudentHandler *http.StudentHandler
}

func (r *RouteConfig) Setup() {
	r.App.Post("/student", r.StudentHandler.Insert)
}
